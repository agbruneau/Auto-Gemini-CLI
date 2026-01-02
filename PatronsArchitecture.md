# 🏗️ Architecture et Design Patterns

Ce document détaille les modèles d'architecture et les choix de conception implémentés dans ce projet.

## 🧩 Patrons d'Architecture

### 1. Event-Driven Architecture (EDA)

Induit un découplage total entre les composants via l'asynchronisme.

- **Implémentation** : Kafka sert de bus de messages.
- **Impact** : Haute disponibilité et extensibilité simplifiée.

### 2. Event Carried State Transfer (ECST)

Chaque message est "autonome" et contient l'intégralité des données nécessaires.

- **Bénéfice** : Pas d'appels API Synchrones vers d'autres services ou bases de données.
- **Fichiers** : [order.go](file:///c:/Users/agbru/OneDrive/Documents/GitHub/PubSubKafka/order.go) définit la structure enrichie.

### 3. Audit Trail & Technical Logging

Séparation des préoccupations en matière de journalisation.

- **Service Monitoring** (`tracker.log`) : Métriques et santé technique.
- **Business Audit** (`tracker.events`) : Journal immuable des flux métier.

### 4. Graceful Shutdown

Les services interceptent les signaux `SIGINT` / `SIGTERM`.

- **Mécanique** : Flush des buffers Kafka et fermeture sécurisée des descripteurs de fichiers.

## 🛠️ Infrastructure & DevOps

- **Kafka mode KRaft** : Suppression de la dépendance à Zookeeper pour plus de simplicité.
- **Go Build Tags** : Gestion des points d'entrée multiples via des tags de compilation (`producer`, `tracker`, `monitor`).
- **Scripts d'orchestration** : [start.sh](file:///c:/Users/agbru/OneDrive/Documents/GitHub/PubSubKafka/start.sh) et [stop.sh](file:///c:/Users/agbru/OneDrive/Documents/GitHub/PubSubKafka/stop.sh) pour une gestion automatisée du cycle de vie.
