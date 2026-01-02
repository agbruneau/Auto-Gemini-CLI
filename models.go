/*
Ce fichier contient les structures de données partagées entre les différents
composants du système (tracker, log_monitor).

Ces structures sont utilisées pour :
- La journalisation structurée (logs système et audit trail)
- La sérialisation/désérialisation des entrées de logs JSON
*/

package main

import "encoding/json"

// LogLevel définit les niveaux de sévérité pour les logs structurés.
type LogLevel string

const (
	LogLevelINFO  LogLevel = "INFO"
	LogLevelERROR LogLevel = "ERROR"
)

// LogEntry est la structure d'un log écrit dans `tracker.log`.
// Elle est conçue pour le patron "Application Health Monitoring".
// Chaque entrée est un log structuré (JSON) contenant des informations sur l'état
// de l'application (démarrage, arrêt, erreurs, métriques). Ce format est optimisé
// pour être ingéré, parsé et visualisé par des outils de monitoring et d'alerte.
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`       // Horodatage du log au format RFC3339.
	Level     LogLevel               `json:"level"`           // Niveau de sévérité (INFO, ERROR).
	Message   string                 `json:"message"`         // Message principal du log.
	Service   string                 `json:"service"`         // Nom du service émetteur.
	Error     string                 `json:"error,omitempty"` // Message d'erreur, si applicable.
	Metadata  map[string]interface{} `json:"metadata,omitempty"` // Données contextuelles supplémentaires.
}

// EventEntry est la structure d'un événement écrit dans `tracker.events`.
// Elle implémente le patron "Audit Trail" en capturant une copie fidèle et immuable
// de chaque message reçu de Kafka, avec ses métadonnées.
//
// Chaque entrée contient le message brut, le résultat de la tentative de désérialisation,
// et des informations contextuelles comme le topic, la partition et l'offset.
// Ce journal est la source de vérité pour l'audit, le rejeu d'événements et le débogage.
type EventEntry struct {
	Timestamp      string          `json:"timestamp"`             // Horodatage de la réception au format RFC3339.
	EventType      string          `json:"event_type"`            // Type d'événement (ex: "message.received").
	KafkaTopic     string          `json:"kafka_topic"`           // Topic Kafka d'origine.
	KafkaPartition int32           `json:"kafka_partition"`       // Partition Kafka d'origine.
	KafkaOffset    int64           `json:"kafka_offset"`          // Offset du message dans la partition.
	RawMessage     string          `json:"raw_message"`           // Contenu brut du message.
	MessageSize    int             `json:"message_size"`          // Taille du message en octets.
	Deserialized   bool            `json:"deserialized"`          // Indique si la désérialisation a réussi.
	Error          string          `json:"error,omitempty"`       // Erreur de désérialisation, si applicable.
	OrderFull      json.RawMessage `json:"order_full,omitempty"`  // Contenu complet de la commande désérialisée.
}

