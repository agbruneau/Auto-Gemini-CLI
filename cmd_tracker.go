//go:build tracker
// +build tracker

/*
Package main provides the entry point for the Tracker (Consumer) service.

The Tracker subscribes to Kafka topics, processes events, and maintains logs.
To compile: go build -tags tracker -o tracker.exe
*/

package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Charger la configuration
	config := NewTrackerConfig()

	// Créer et initialiser le tracker
	tracker := NewTracker(config)
	if err := tracker.Initialize(); err != nil {
		log.Fatalf("Erreur fatale lors de l'initialisation: %v", err)
	}
	defer tracker.Close()

	fmt.Println("🟢 Le consommateur est en cours d'exécution...")
	fmt.Printf("📝 Logs d'observabilité système dans %s\n", config.LogFile)
	fmt.Printf("📋 Journalisation complète des messages dans %s\n", config.EventsFile)

	// Gestion des signaux d'arrêt
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Lancer le tracker dans une goroutine
	done := make(chan struct{})
	go func() {
		tracker.Run()
		close(done)
	}()

	// Attendre un signal d'arrêt
	<-sigchan
	fmt.Println("\n⚠️ Signal d'arrêt reçu...")
	tracker.Stop()
	<-done

	fmt.Println("🔴 Le consommateur est arrêté.")
}
