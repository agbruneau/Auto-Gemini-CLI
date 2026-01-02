//go:build producer
// +build producer

/*
Package main provides the entry point for the Producer service.

The Producer simulates order generation and streams events to Apache Kafka.
To compile: go build -tags producer -o producer.exe
*/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Charger la configuration
	config := NewProducerConfig()

	// Créer et initialiser le producteur
	producer := NewOrderProducer(config)
	if err := producer.Initialize(); err != nil {
		fmt.Printf("Erreur fatale lors de l'initialisation: %v\n", err)
		os.Exit(1)
	}
	defer producer.Close()

	fmt.Println("🟢 Le producteur est démarré et prêt à envoyer des messages...")
	fmt.Printf("📤 Publication sur le topic '%s'\n", config.Topic)

	// Gestion des signaux d'arrêt
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Lancer la boucle de production
	producer.Run(sigchan)
}
