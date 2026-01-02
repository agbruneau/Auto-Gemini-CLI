# 🚀 Système de Suivi de Commandes Kafka (Kafka Order Tracking)

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Kafka](https://img.shields.io/badge/Apache_Kafka-3.7.0-white?style=flat&logo=apache-kafka)](https://kafka.apache.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

Bienvenue dans le projet **Kafka Order Tracking**. Cette implémentation de référence en **Go** illustre une architecture orientée événements (EDA) moderne et robuste utilisant **Apache Kafka**.

Le système simule un flux de commandes e-commerce complet — de la génération à la consommation — tout en offrant une observabilité avancée via une interface terminal (TUI) interactive.

---

## 🏗 Architecture du Système

Le projet repose sur trois services principaux totalement découplés :

1.  **📦 Producteur (`producer`)** : Génère des flux de commandes aléatoires enrichies (simulant des achats clients) et les publie dans le topic Kafka `orders`.
2.  **⚙️ Consommateur (`tracker`)** : S'abonne au topic `orders`, traite les messages en temps réel et maintient une piste d'audit exhaustive.
3.  **📊 Moniteur (`log_monitor`)** : Une interface graphique en terminal (TUI) offrant une visualisation en temps réel des métriques de performance (débit, latence, succès) et des logs système.

---

## 🌟 Principes et Design Patterns

Ce projet met en œuvre les standards industriels pour les systèmes distribués :

- **Event-Driven Architecture (EDA)** : Découplage maximal entre émetteurs et récepteurs.
- **Event Carried State Transfer (ECST)** : Les messages incluent tout le contexte nécessaire (produit, client, prix), rendant les consommateurs autonomes.
- **Guaranteed Delivery (At-Least-Once)** : Utilisation des rapports de livraison (ACK) pour garantir l'intégrité des données.
- **Observabilité Duale** :
  - **Health Monitoring** (`tracker.log`) : Logs techniques structurés (JSON) pour le monitoring.
  - **Audit Trail** (`tracker.events`) : Journal immuable de tous les événements métier reçus.
- **Graceful Shutdown** : Gestion rigoureuse des signaux système (SIGTERM, SIGINT) pour un arrêt sans perte de données.
- **Idempotence Opérationnelle** : Automatisation de la création des ressources Kafka via des scripts robustes.

---

## 🛠 Prérequis

Avant de commencer, assurez-vous d'avoir installé :

1.  **Docker** et **Docker Compose** (V2).
2.  **Go** (version 1.22 ou supérieure).
3.  Un terminal compatible ANSI (pour le moniteur).
4.  Privilèges `sudo` (requis pour les commandes Docker dans les scripts).

---

## 🚀 Démarrage Rapide

Le projet fournit un automate d'orchestration pour déployer l'environnement complet.

1.  **Initialisez l'infrastructure et lancez les services** :
    ```bash
    ./start.sh
    ```

**Actions réalisées par le script :**

- Déploiement du cluster Kafka (mode KRaft) via Docker Compose.
- Vérification de la disponibilité du broker.
- Création idempotente du topic `orders`.
- Lancement des services Go (**Producer** et **Tracker**) en arrière-plan.

---

## 📊 Utilisation et Monitoring

Une fois le système lancé, plusieurs méthodes s'offrent à vous pour observer l'activité.

### 1. Le Moniteur Interactif (Recommandé)

Pour une vue d'ensemble visuelle (Tableau de bord, graphiques, logs défilants), lancez le moniteur dans un **nouveau terminal** :

```bash
go run -tags monitor cmd_monitor.go log_monitor.go models.go constants.go
```

- **Touches** : `q` ou `Ctrl+C` pour quitter.
- **Fonctionnalités** : Affiche le débit (msg/sec), le taux de succès, et les derniers logs.

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

- **Points d'entrée (`cmd_*.go`)** :
  - `cmd_producer.go` : `main()` du producteur.
  - `cmd_tracker.go` : `main()` du consommateur.
  - `cmd_monitor.go` : `main()` du moniteur TUI.
- **Logique Métier** :
  - `producer.go` : Implémentation de l'envoi Kafka.
  - `tracker.go` : Logique de traitement des messages.
  - `log_monitor.go` : Logique d'affichage TUI.
- **Données partagées** :
  - `models.go` : Structures de données (Logs, Métriques).
  - `order.go` : Définition de la structure `Order`.
  - `constants.go` : Configuration globale (Topics, Fichiers, Timeouts).
- **Scripts** :
  - `start.sh` / `stop.sh` : Gestion du cycle de vie.
  - `docker-compose.yaml` : Infrastructure.

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
