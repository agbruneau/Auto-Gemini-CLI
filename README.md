# Système de Suivi de Commandes Kafka (Kafka Order Tracking System)

Bienvenue dans le projet de démonstration **Kafka Order Tracking**. Ce projet est une implémentation de référence en **Go** illustrant une architecture événementielle (EDA) robuste utilisant **Apache Kafka**. Il simule un flux de commandes e-commerce complet, de la production à la consommation, avec une observabilité avancée.

## 📋 Table des Matières

- [Architecture](#-architecture)
- [Fonctionnalités et Patterns](#-fonctionnalités-et-patterns)
- [Prérequis](#-prérequis)
- [Démarrage Rapide](#-démarrage-rapide)
- [Utilisation et Monitoring](#-utilisation-et-monitoring)
- [Arrêt du Système](#-arrêt-du-système)
- [Structure du Projet](#-structure-du-projet)
- [Développement et Tests](#-développement-et-tests)

---

## 🏗 Architecture

Le système est composé de trois services principaux découplés, communiquant via Kafka ou observant l'état du système via des logs.

```mermaid
graph LR
    P[Producteur (Producer)] -->|Envoie 'Order'| K{Kafka Topic: orders}
    K -->|Consomme 'Order'| T[Consommateur (Tracker)]
    T -->|Écrit| L1[tracker.log (Santé)]
    T -->|Écrit| L2[tracker.events (Audit)]
    M[Moniteur (TUI)] -.->|Lit| L1
    M -.->|Lit| L2
```

1.  **Producteur (`producer`)** : Génère des commandes aléatoires (simulant des achats clients) et les envoie dans le topic Kafka `orders`.
2.  **Consommateur (`tracker`)** : Écoute le topic `orders`, traite les commandes reçues et enregistre le résultat.
3.  **Moniteur (`log_monitor`)** : Une interface graphique en terminal (TUI) qui visualise en temps réel les métriques de performance et les logs.

---

## 🌟 Fonctionnalités et Patterns

Ce projet met en œuvre les meilleures pratiques de l'ingénierie logicielle distribuée :

*   **Event-Driven Architecture (EDA)** : Découplage total entre le producteur et le consommateur.
*   **Event Carried State Transfer (ECST)** : Les messages contiennent tout le contexte nécessaire (produit, client, prix), rendant le consommateur autonome (pas d'appels API externes nécessaires).
*   **Guaranteed Delivery** : Le producteur attend l'accusé de réception (ACK) du broker Kafka pour confirmer l'envoi.
*   **Idempotence** : Le script de démarrage assure que les ressources (topics) ne sont créées que si elles n'existent pas.
*   **Observabilité Duale** :
    *   `tracker.log` : Logs structurés (JSON) pour la santé technique (erreurs, latence).
    *   `tracker.events` : Piste d'audit immuable de tous les événements métier reçus.
*   **Graceful Shutdown** : Gestion propre des signaux (SIGTERM, SIGINT) pour terminer les processus sans perte de données (flush des messages, fermeture des fichiers).

---

## 🛠 Prérequis

Avant de commencer, assurez-vous d'avoir installé :

1.  **Docker** et **Docker Compose** (V2).
2.  **Go** (version 1.22 ou supérieure).
3.  Un terminal compatible ANSI (pour le moniteur).
4.  Privilèges `sudo` (requis pour les commandes Docker dans les scripts).

---

## 🚀 Démarrage Rapide

Le projet fournit un script d'orchestration pour lancer l'environnement complet en une seule commande.

1.  Placez-vous à la racine du projet.
2.  Lancez le script de démarrage :

```bash
./start.sh
```

**Ce que fait le script :**
*   Démarre le conteneur Kafka via Docker Compose.
*   Attend activement que Kafka soit prêt.
*   Crée le topic `orders` de manière idempotente.
*   Lance le **Tracker** (consommateur) en arrière-plan.
*   Lance le **Producer** (producteur) en arrière-plan (mais attache le script à son processus).

---

## 📊 Utilisation et Monitoring

Une fois le système lancé, plusieurs méthodes s'offrent à vous pour observer l'activité.

### 1. Le Moniteur Interactif (Recommandé)

Pour une vue d'ensemble visuelle (Tableau de bord, graphiques, logs défilants), lancez le moniteur dans un **nouveau terminal** :

```bash
go run -tags monitor cmd_monitor.go log_monitor.go models.go constants.go
```

*   **Touches** : `q` ou `Ctrl+C` pour quitter.
*   **Fonctionnalités** : Affiche le débit (msg/sec), le taux de succès, et les derniers logs.

### 2. Observation des Logs Bruts

Vous pouvez suivre les fichiers de logs générés en temps réel :

```bash
# Pour voir l'activité métier (Audit)
tail -f tracker.events

# Pour voir la santé technique (Logs JSON)
# (Si vous avez 'jq' installé pour le formatage)
tail -f tracker.log | jq
```

---

## 🛑 Arrêt du Système

Pour arrêter proprement tous les composants (processus Go et conteneurs Docker), utilisez le script dédié :

```bash
./stop.sh
```

Ce script utilise les fichiers PID (`producer.pid`, `tracker.pid`) pour envoyer des signaux de terminaison (SIGTERM) aux processus Go, leur laissant le temps de finir leur travail en cours, avant d'arrêter l'infrastructure Docker.

---

## 📂 Structure du Projet

L'organisation des fichiers suit une logique modulaire :

*   **Points d'entrée (`cmd_*.go`)** :
    *   `cmd_producer.go` : `main()` du producteur.
    *   `cmd_tracker.go` : `main()` du consommateur.
    *   `cmd_monitor.go` : `main()` du moniteur TUI.
*   **Logique Métier** :
    *   `producer.go` : Implémentation de l'envoi Kafka.
    *   `tracker.go` : Logique de traitement des messages.
    *   `log_monitor.go` : Logique d'affichage TUI.
*   **Données partagées** :
    *   `models.go` : Structures de données (Logs, Métriques).
    *   `order.go` : Définition de la structure `Order`.
    *   `constants.go` : Configuration globale (Topics, Fichiers, Timeouts).
*   **Scripts** :
    *   `start.sh` / `stop.sh` : Gestion du cycle de vie.
    *   `docker-compose.yaml` : Infrastructure.

---

## 💻 Développement et Tests

Ce projet utilise des **Build Tags** Go (`producer`, `tracker`, `monitor`, `kafka`) pour compiler conditionnellement les différents composants.

### Compilation Manuelle

Si vous ne souhaitez pas utiliser `go run`, vous pouvez compiler les binaires :

```bash
# Compiler le Producteur
go build -tags producer -o producer cmd_producer.go producer.go order.go models.go constants.go

# Compiler le Tracker
go build -tags tracker -o tracker cmd_tracker.go tracker.go order.go models.go constants.go

# Compiler le Moniteur
go build -tags monitor -o monitor cmd_monitor.go log_monitor.go models.go constants.go
```

### Exécution des Tests

Les tests unitaires nécessitent également les tags appropriés et les fichiers dépendants :

```bash
# Tester la logique du Producteur
go test -tags kafka,producer producer.go producer_test.go order.go constants.go

# Tester la logique du Tracker
go test -tags kafka,tracker tracker.go tracker_test.go order.go constants.go models.go

# Tester le Moniteur
go test -tags monitor log_monitor.go log_monitor_test.go models.go constants.go
```
