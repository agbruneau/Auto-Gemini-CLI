//go:build kafka
// +build kafka

/*
tracker.go is the primary consumer for the Kafka Order Tracking System.
It subscribes to the 'orders' topic, processes incoming JSON events, and
maintains high system observability.

Architecture Highlights:
- **Event Consumption**: Implements a robust consumer loop for the 'orders' topic.
- **Advanced Observability**: Uses a dual-logging strategy:
  1. **Health Monitoring** (tracker.log): Structured system metrics and lifecycle events.
  2. **Audit Trail** (tracker.events): Immutable journal of every received message.
- **Operational Metrics**: Periodically reports throughput and success rates.
- **Safe Shutdown**: Leverages context and signal handling for zero-data-loss stop.
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Note: Diagnostic and Audit models are defined in models.go.
// Note: Configuration constants are defined in constants.go for project-wide consistency.

// TrackerConfig encapsulates the tracker service's operating parameters.
// This structure supports overrides via environment variables for cloud-native compatibility.
type TrackerConfig struct {
	KafkaBroker     string        // Adresse du broker Kafka
	ConsumerGroup   string        // Groupe de consommateurs Kafka
	Topic           string        // Topic Kafka à consommer
	LogFile         string        // Fichier de logs système
	EventsFile      string        // Fichier de journal d'audit
	MetricsInterval time.Duration // Intervalle entre les métriques périodiques
	ReadTimeout     time.Duration // Timeout de lecture des messages
	MaxErrors       int           // Nombre maximum d'erreurs consécutives
}

// NewTrackerConfig initializes a configuration object with default values,
// subsequently applying any environment variable overrides if present.
func NewTrackerConfig() *TrackerConfig {
	config := &TrackerConfig{
		KafkaBroker:     DefaultKafkaBroker,
		ConsumerGroup:   DefaultConsumerGroup,
		Topic:           DefaultTopic,
		LogFile:         TrackerLogFile,
		EventsFile:      TrackerEventsFile,
		MetricsInterval: TrackerMetricsInterval,
		ReadTimeout:     TrackerConsumerReadTimeout,
		MaxErrors:       TrackerMaxConsecutiveErrors,
	}

	// Surcharge depuis les variables d'environnement
	if broker := os.Getenv("KAFKA_BROKER"); broker != "" {
		config.KafkaBroker = broker
	}
	if group := os.Getenv("KAFKA_CONSUMER_GROUP"); group != "" {
		config.ConsumerGroup = group
	}
	if topic := os.Getenv("KAFKA_TOPIC"); topic != "" {
		config.Topic = topic
	}

	return config
}

// Logger provides synchronized, high-performance encoding for structured logs.
type Logger struct {
	file    *os.File
	encoder *json.Encoder
	mu      sync.Mutex
}

// SystemMetrics aggregates performance counters for the consumer service.
// Access is protected by a RWMutex to ensure thread-safety during concurrent processing.
type SystemMetrics struct {
	mu                sync.RWMutex
	StartTime         time.Time
	MessagesReceived  int64
	MessagesProcessed int64
	MessagesFailed    int64
	LastMessageTime   time.Time
}

// Tracker is the primary service managing Kafka message consumption and observability.
// It orchestrates dual-stream logging, metrics collection, and graceful shutdown.
type Tracker struct {
	config      *TrackerConfig
	logLogger   *Logger
	eventLogger *Logger
	metrics     *SystemMetrics
	consumer    *kafka.Consumer
	stopChan    chan struct{}
	running     bool
	mu          sync.Mutex
}

// NewTracker creates a new instance of the Tracker service.
func NewTracker(config *TrackerConfig) *Tracker {
	return &Tracker{
		config:   config,
		metrics:  &SystemMetrics{StartTime: time.Now()},
		stopChan: make(chan struct{}),
	}
}

// Initialize sets up the dual-stream loggers and initializes the Kafka consumer.
func (t *Tracker) Initialize() error {
	var err error

	// Initialiser les loggers
	t.logLogger, err = newLogger(t.config.LogFile)
	if err != nil {
		return fmt.Errorf("impossible d'initialiser le logger système: %w", err)
	}

	t.eventLogger, err = newLogger(t.config.EventsFile)
	if err != nil {
		t.logLogger.Close()
		return fmt.Errorf("impossible d'initialiser le logger d'événements: %w", err)
	}

	t.logLogger.Log(LogLevelINFO, "Système de journalisation initialisé", map[string]interface{}{
		"log_file":    t.config.LogFile,
		"events_file": t.config.EventsFile,
	})

	// Initialiser le consommateur Kafka
	t.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": t.config.KafkaBroker,
		"group.id":          t.config.ConsumerGroup,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		t.logLogger.LogError("Erreur lors de la création du consommateur", err, nil)
		t.Close()
		return fmt.Errorf("impossible de créer le consommateur Kafka: %w", err)
	}

	// S'abonner au topic
	err = t.consumer.SubscribeTopics([]string{t.config.Topic}, nil)
	if err != nil {
		t.logLogger.LogError("Erreur lors de l'abonnement au topic", err, map[string]interface{}{"topic": t.config.Topic})
		t.Close()
		return fmt.Errorf("impossible de s'abonner au topic: %w", err)
	}

	t.logLogger.Log(LogLevelINFO, "Consommateur démarré et abonné au topic '"+t.config.Topic+"'", nil)
	return nil
}

// Run enters the main consumption loop, processing messages until a cessation is signaled.
func (t *Tracker) Run() {
	t.mu.Lock()
	t.running = true
	t.mu.Unlock()

	// Démarrer les métriques périodiques
	go t.logPeriodicMetrics()

	consecutiveErrors := 0

	for t.isRunning() {
		msg, err := t.consumer.ReadMessage(t.config.ReadTimeout)
		if err != nil {
			shouldStop := t.handleKafkaError(err, &consecutiveErrors)
			if shouldStop {
				break
			}
			continue
		}

		consecutiveErrors = 0
		t.processMessage(msg)
	}
}

// isRunning returns true if the tracker service is currently active.
func (t *Tracker) isRunning() bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	return t.running
}

// handleKafkaError manages connectivity and protocol errors from the Kafka broker.
// Returns true if the error is terminal and the service should halt.
func (t *Tracker) handleKafkaError(err error, consecutiveErrors *int) bool {
	kafkaErr, ok := err.(kafka.Error)
	if !ok {
		return false
	}

	// Timeout normal, pas une erreur
	if kafkaErr.Code() == kafka.ErrTimedOut {
		*consecutiveErrors = 0
		return false
	}

	// Vérifier si c'est une erreur de connexion critique
	errorMsg := err.Error()
	isShutdownError := strings.Contains(errorMsg, "brokers are down") ||
		strings.Contains(errorMsg, "Connection refused") ||
		kafkaErr.Code() == kafka.ErrAllBrokersDown

	if isShutdownError {
		*consecutiveErrors++
		if *consecutiveErrors >= t.config.MaxErrors {
			t.logLogger.Log(LogLevelINFO, "Kafka semble être arrêté, arrêt du consommateur", map[string]interface{}{
				"consecutive_errors": *consecutiveErrors,
				"reason":             "brokers_unavailable",
			})
			return true
		}
		return false
	}

	// Autres erreurs
	t.logLogger.LogError("Erreur de lecture du message Kafka", err, nil)
	*consecutiveErrors++
	if *consecutiveErrors >= t.config.MaxErrors {
		t.logLogger.LogError("Trop d'erreurs consécutives, arrêt du consommateur", err, map[string]interface{}{
			"consecutive_errors": *consecutiveErrors,
		})
		return true
	}

	return false
}

// processMessage handles the deserialization, logging, and metrics for a single message.
func (t *Tracker) processMessage(msg *kafka.Message) {
	var order Order
	deserializationErr := json.Unmarshal(msg.Value, &order)

	// Journaliser l'événement (toujours)
	var orderForLog *Order
	if deserializationErr == nil {
		orderForLog = &order
	}
	t.eventLogger.LogEvent(msg, orderForLog, deserializationErr)

	// Mettre à jour les métriques et traiter le message
	if deserializationErr != nil {
		t.metrics.recordMetrics(false, true)
		t.logLogger.LogError("Erreur de désérialisation du message", deserializationErr, map[string]interface{}{
			"kafka_offset": msg.TopicPartition.Offset,
			"raw_message":  string(msg.Value),
		})
	} else {
		t.metrics.recordMetrics(true, false)
		displayOrder(&order)
	}
}

// logPeriodicMetrics periodically exports system performance to the diagnostics log.
func (t *Tracker) logPeriodicMetrics() {
	ticker := time.NewTicker(t.config.MetricsInterval)
	defer ticker.Stop()

	for {
		select {
		case <-t.stopChan:
			return
		case <-ticker.C:
			t.metrics.mu.RLock()
			uptime := time.Since(t.metrics.StartTime)
			var successRate float64
			if t.metrics.MessagesReceived > 0 {
				successRate = float64(t.metrics.MessagesProcessed) / float64(t.metrics.MessagesReceived) * 100
			}
			var messagesPerSecond float64
			if uptime.Seconds() > 0 {
				messagesPerSecond = float64(t.metrics.MessagesReceived) / uptime.Seconds()
			}
			t.metrics.mu.RUnlock()

			t.logLogger.Log(LogLevelINFO, "Métriques système périodiques", map[string]interface{}{
				"uptime_seconds":       uptime.Seconds(),
				"messages_received":    t.metrics.MessagesReceived,
				"messages_processed":   t.metrics.MessagesProcessed,
				"messages_failed":      t.metrics.MessagesFailed,
				"success_rate_percent": fmt.Sprintf("%.2f", successRate),
				"messages_per_second":  fmt.Sprintf("%.2f", messagesPerSecond),
			})
		}
	}
}

// Stop triggers a graceful cessation of the consumption loop.
func (t *Tracker) Stop() {
	t.mu.Lock()
	t.running = false
	t.mu.Unlock()

	close(t.stopChan)

	// Log final
	uptime := time.Since(t.metrics.StartTime)
	t.logLogger.Log(LogLevelINFO, "Consommateur arrêté proprement", map[string]interface{}{
		"uptime_seconds":           uptime.Seconds(),
		"total_messages_received":  t.metrics.MessagesReceived,
		"total_messages_processed": t.metrics.MessagesProcessed,
		"total_messages_failed":    t.metrics.MessagesFailed,
	})
}

// Close releases all underlying resources, including the Kafka consumer and loggers.
func (t *Tracker) Close() {
	if t.consumer != nil {
		t.consumer.Close()
	}
	if t.logLogger != nil {
		t.logLogger.Close()
	}
	if t.eventLogger != nil {
		t.eventLogger.Close()
	}
}

// Global variables for legacy and test compatibility.
var (
	logLogger     *Logger
	eventLogger   *Logger
	systemMetrics = &SystemMetrics{StartTime: time.Now()}
)

// newLogger initializes a new Logger specialized for structured JSON output.
func newLogger(filename string) (*Logger, error) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("impossible d'ouvrir le fichier %s: %v", filename, err)
	}
	return &Logger{
		file:    file,
		encoder: json.NewEncoder(file),
	}, nil
}

// Log writes a structured entry to the diagnostic record (tracker.log).
func (l *Logger) Log(level LogLevel, message string, metadata map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     level,
		Message:   message,
		Service:   TrackerServiceName,
		Metadata:  metadata,
	}
	if err := l.encoder.Encode(entry); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur d'encodage du log: %v\n", err)
	}
}

// LogError is a convenience wrapper for recording system anomalies in the diagnostic record.
func (l *Logger) LogError(message string, err error, metadata map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if metadata == nil {
		metadata = make(map[string]interface{})
	}
	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Level:     LogLevelERROR,
		Message:   message,
		Service:   TrackerServiceName,
		Error:     err.Error(),
		Metadata:  metadata,
	}
	if encodeErr := l.encoder.Encode(entry); encodeErr != nil {
		fmt.Fprintf(os.Stderr, "Erreur d'encodage du log d'erreur: %v\n", encodeErr)
	}
}

// LogEvent records a high-fidelity audit entry for every incoming message.
// This function is the core of the "Audit Trail" implementation, capturing
// messages regardless of validity to ensure full traceability.
func (l *Logger) LogEvent(msg *kafka.Message, order *Order, deserializationError error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	eventType := "message.received"
	deserialized := order != nil

	if deserializationError != nil {
		eventType = "message.received.deserialization_error"
	}

	event := EventEntry{
		Timestamp:      time.Now().UTC().Format(time.RFC3339),
		EventType:      eventType,
		KafkaTopic:     *msg.TopicPartition.Topic,
		KafkaPartition: msg.TopicPartition.Partition,
		KafkaOffset:    int64(msg.TopicPartition.Offset),
		RawMessage:     string(msg.Value),
		MessageSize:    len(msg.Value),
		Deserialized:   deserialized,
	}

	if deserialized {
		orderJSON, marshalErr := json.Marshal(order)
		if marshalErr != nil {
			fmt.Fprintf(os.Stderr, "Erreur de sérialisation de la commande: %v\n", marshalErr)
		} else {
			event.OrderFull = json.RawMessage(orderJSON)
		}
	}

	if deserializationError != nil {
		event.Error = deserializationError.Error()
	}

	if err := l.encoder.Encode(event); err != nil {
		fmt.Fprintf(os.Stderr, "Erreur d'encodage de l'événement: %v\n", err)
	}
}

// Close ensures all file handles are properly released.
func (l *Logger) Close() {
	if l != nil && l.file != nil {
		if err := l.file.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "Erreur lors de la fermeture du fichier de log: %v\n", err)
		}
	}
}

// recordMetrics atomiquement updates the performance counters.
func (sm *SystemMetrics) recordMetrics(processed, failed bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	sm.MessagesReceived++
	if processed {
		sm.MessagesProcessed++
	}
	if failed {
		sm.MessagesFailed++
	}
	sm.LastMessageTime = time.Now()
}

// Note: La fonction main() est définie dans cmd_tracker.go avec le build tag "tracker"
// Pour compiler: go build -tags tracker -o tracker.exe

// displayOrder prints a formatted summary of an order to the console for real-time monitoring.
func displayOrder(order *Order) {
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Printf("📦 COMMANDE REÇUE #%d (ID: %s)\n", order.Sequence, order.OrderID)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("Client: %s (%s)\n", order.CustomerInfo.Name, order.CustomerInfo.CustomerID)
	fmt.Printf("Statut: %s | Total: %.2f %s\n", order.Status, order.Total, order.Currency)
	fmt.Println("Articles:")
	for _, item := range order.Items {
		fmt.Printf("  - %s (x%d) @ %.2f %s\n", item.ItemName, item.Quantity, item.UnitPrice, order.Currency)
	}
	fmt.Println(strings.Repeat("=", 80))
}
