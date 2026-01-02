//go:build kafka
// +build kafka

/*
Ce programme Go (`producer.go`) agit comme un producteur de messages pour Apache Kafka.
Son rôle est de simuler la création de commandes enrichies et de les envoyer
de manière continue à un topic Kafka.

Il implémente une logique de production robuste qui met en œuvre plusieurs bonnes pratiques :
- **Event Carried State Transfer** : Il génère des données de commande complètes et autonomes.
- **Publisher/Subscriber** : Il publie des messages dans un topic Kafka.
- **Guaranteed Delivery** : Il utilise un canal de rapport de livraison (`deliveryReport`)
  pour s'assurer que chaque message est bien reçu par le broker Kafka.
- **Graceful Shutdown** : Il intercepte les signaux d'arrêt du système pour terminer proprement
  son exécution, en s'assurant que tous les messages en attente dans le tampon sont
  envoyés avant de quitter (`producer.Flush`).
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
)

// Note: Les constantes sont définies dans constants.go pour éviter les duplications.

// ProducerConfig contient la configuration du service producteur.
// Elle peut être chargée depuis des variables d'environnement.
type ProducerConfig struct {
	KafkaBroker     string        // Adresse du broker Kafka
	Topic           string        // Topic Kafka pour la publication
	MessageInterval time.Duration // Intervalle entre les messages
	FlushTimeout    int           // Timeout en ms pour le flush final
	TaxRate         float64       // Taux de taxe à appliquer
	ShippingFee     float64       // Frais de livraison
	Currency        string        // Devise par défaut
	PaymentMethod   string        // Méthode de paiement par défaut
	Warehouse       string        // Entrepôt par défaut
}

// NewProducerConfig crée une configuration avec les valeurs par défaut,
// surchargées par les variables d'environnement si elles sont définies.
func NewProducerConfig() *ProducerConfig {
	config := &ProducerConfig{
		KafkaBroker:     DefaultKafkaBroker,
		Topic:           DefaultTopic,
		MessageInterval: ProducerMessageInterval,
		FlushTimeout:    FlushTimeoutMs,
		TaxRate:         ProducerDefaultTaxRate,
		ShippingFee:     ProducerDefaultShippingFee,
		Currency:        ProducerDefaultCurrency,
		PaymentMethod:   ProducerDefaultPayment,
		Warehouse:       ProducerDefaultWarehouse,
	}

	// Surcharge depuis les variables d'environnement
	if broker := os.Getenv("KAFKA_BROKER"); broker != "" {
		config.KafkaBroker = broker
	}
	if topic := os.Getenv("KAFKA_TOPIC"); topic != "" {
		config.Topic = topic
	}

	return config
}

// OrderTemplate définit un modèle pour générer des commandes de test.
type OrderTemplate struct {
	User     string  // Identifiant du client
	Item     string  // Nom de l'article
	Quantity int     // Quantité commandée
	Price    float64 // Prix unitaire
}

// DefaultOrderTemplates contient les templates de commandes par défaut.
var DefaultOrderTemplates = []OrderTemplate{
	{User: "client01", Item: "espresso", Quantity: 2, Price: 2.50},
	{User: "client02", Item: "cappuccino", Quantity: 3, Price: 3.20},
	{User: "client03", Item: "latte", Quantity: 4, Price: 3.50},
	{User: "client04", Item: "macchiato", Quantity: 5, Price: 3.00},
	{User: "client05", Item: "flat white", Quantity: 6, Price: 3.30},
	{User: "client06", Item: "mocha", Quantity: 7, Price: 4.00},
	{User: "client07", Item: "americano", Quantity: 8, Price: 2.80},
	{User: "client08", Item: "chai latte", Quantity: 9, Price: 3.80},
	{User: "client09", Item: "matcha", Quantity: 10, Price: 4.50},
	{User: "client10", Item: "smoothie fraise", Quantity: 11, Price: 5.50},
}

// OrderProducer est le service principal qui gère la production de messages Kafka.
// Il encapsule le producteur Kafka, la configuration et les templates de commandes.
type OrderProducer struct {
	config       *ProducerConfig
	producer     *kafka.Producer
	deliveryChan chan kafka.Event
	templates    []OrderTemplate
	sequence     int
	running      bool
}

// NewOrderProducer crée une nouvelle instance du service OrderProducer.
func NewOrderProducer(config *ProducerConfig) *OrderProducer {
	return &OrderProducer{
		config:    config,
		templates: DefaultOrderTemplates,
		sequence:  1,
	}
}

// Initialize initialise le producteur Kafka.
func (p *OrderProducer) Initialize() error {
	var err error
	p.producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": p.config.KafkaBroker,
	})
	if err != nil {
		return fmt.Errorf("impossible de créer le producteur Kafka: %w", err)
	}

	p.deliveryChan = make(chan kafka.Event, ProducerDeliveryChannelSize)
	go p.handleDeliveryReports()

	return nil
}

// handleDeliveryReports traite les rapports de livraison dans une goroutine dédiée.
func (p *OrderProducer) handleDeliveryReports() {
	for e := range p.deliveryChan {
		m := e.(*kafka.Message)
		if m.TopicPartition.Error != nil {
			fmt.Printf("❌ La livraison du message a échoué: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("✅ Message livré avec succès au topic %s (partition %d) à l'offset %d\n",
				*m.TopicPartition.Topic,
				m.TopicPartition.Partition,
				m.TopicPartition.Offset)
		}
	}
}

// GenerateOrder crée une commande enrichie à partir d'un template et d'un numéro de séquence.
func (p *OrderProducer) GenerateOrder(template OrderTemplate, sequence int) Order {
	// Calculs financiers
	itemTotal := float64(template.Quantity) * template.Price
	tax := itemTotal * p.config.TaxRate
	total := itemTotal + tax + p.config.ShippingFee

	return Order{
		OrderID:  uuid.New().String(),
		Sequence: sequence,
		Status:   "pending",
		Items: []OrderItem{
			{
				ItemID:     fmt.Sprintf("item-%s", template.Item),
				ItemName:   template.Item,
				Quantity:   template.Quantity,
				UnitPrice:  template.Price,
				TotalPrice: itemTotal,
			},
		},
		SubTotal:        itemTotal,
		Tax:             tax,
		ShippingFee:     p.config.ShippingFee,
		Total:           total,
		Currency:        p.config.Currency,
		PaymentMethod:   p.config.PaymentMethod,
		ShippingAddress: fmt.Sprintf("%d Rue de la Paix, 75000 Paris", sequence),
		Metadata: OrderMetadata{
			Timestamp:     time.Now().UTC().Format(time.RFC3339),
			Version:       "1.1",
			EventType:     "order.created",
			Source:        "producer-service",
			CorrelationID: uuid.New().String(),
		},
		CustomerInfo: CustomerInfo{
			CustomerID:   template.User,
			Name:         fmt.Sprintf("Client %s", template.User),
			Email:        fmt.Sprintf("%s@example.com", template.User),
			Phone:        "+33 6 00 00 00 00",
			Address:      fmt.Sprintf("%d Rue de la Paix, 75000 Paris", sequence),
			LoyaltyLevel: "silver",
		},
		InventoryStatus: []InventoryStatus{
			{
				ItemID:       fmt.Sprintf("item-%s", template.Item),
				ItemName:     template.Item,
				AvailableQty: 100 - template.Quantity,
				ReservedQty:  template.Quantity,
				UnitPrice:    template.Price,
				InStock:      true,
				Warehouse:    p.config.Warehouse,
			},
		},
	}
}

// ProduceOrder génère et envoie une commande au topic Kafka.
func (p *OrderProducer) ProduceOrder() error {
	template := p.templates[p.sequence%len(p.templates)]
	order := p.GenerateOrder(template, p.sequence)

	value, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("erreur de sérialisation JSON: %w", err)
	}

	topic := p.config.Topic
	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, p.deliveryChan)

	if err != nil {
		return fmt.Errorf("erreur lors de la production du message: %w", err)
	}

	p.sequence++
	return nil
}

// Run démarre la boucle de production de messages.
func (p *OrderProducer) Run(stopChan <-chan os.Signal) {
	p.running = true
	for p.running {
		select {
		case <-stopChan:
			fmt.Println("\n⚠️  Signal d'arrêt reçu. Fin de la production de nouveaux messages...")
			p.running = false
		default:
			if err := p.ProduceOrder(); err != nil {
				fmt.Printf("Erreur: %v\n", err)
			}
			time.Sleep(p.config.MessageInterval)
		}
	}
}

// Close ferme proprement le producteur et envoie les messages en attente.
func (p *OrderProducer) Close() {
	fmt.Println("⏳ Envoi des messages restants en file d'attente...")
	remainingMessages := p.producer.Flush(p.config.FlushTimeout)
	if remainingMessages > 0 {
		fmt.Printf("⚠️  %d messages n'ont pas pu être envoyés.\n", remainingMessages)
	} else {
		fmt.Println("✅ Tous les messages ont été envoyés avec succès.")
	}
	p.producer.Close()
}

// Note: La fonction main() est définie dans cmd_producer.go avec le build tag "producer"
// Pour compiler: go build -tags producer -o producer.exe