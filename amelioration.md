# 📋 Plan d'Améliorations Priorisé - PubSub Kafka Demo

Ce document recense et priorise les améliorations techniques pour faire évoluer le projet d'une démonstration vers une application robuste prête pour la production.

## 🏆 Priorisation & Roadmap

La priorisation est basée sur l'impact (stabilité, maintenabilité) par rapport à l'effort.

| Priorité | Domaine | Amélioration Clé | Impact |
|---|---|---|---|
| **🔴 Critique** | **Architecture** | **1.1 Structure de packages Standard** | Fondamental pour la maintenabilité et les tests. |
| **🔴 Critique** | **Config** | **2.1 Configuration Externe** | Indispensable pour déployer dans différents environnements sans recompiler. |
| **🔴 Critique** | **Fiabilité** | **6.1 Retry Pattern** | Nécessaire pour gérer les pannes réseaux transitoires. |
| **🟠 Élevée** | **Tests** | **4.2 Couverture de tests** | Sécurise les refactorings futurs. |
| **🟠 Élevée** | **DevOps** | **7.1 Docker Multi-stage** | Optimise la taille des images et la sécurité pour la prod. |
| **🟠 Élevée** | **CI/CD** | **11.1 GitHub Actions** | Automatise la qualité du code. |
| **🟡 Moyenne** | **Observabilité** | **5.2 Métriques Prometheus** | Standard de l'industrie (remplace le `log_monitor` custom à terme). |
| **🟡 Moyenne** | **Sécurité** | **3.1 Auth Kafka** | Critique en prod, mais optionnel en local/demo. |
| **🟢 Basse** | **Fonctionnalité** | **8.1 Multi-topics / 9.2 Web UI** | Extensions fonctionnelles non bloquantes. |

---

## 1. 🏗️ Architecture et Organisation du Code (Critique)

### 1.1 Migration vers une structure de packages Go standard
**Priorité : Critique**
Actuellement, tout est dans le package `main`. Cela empêche les tests unitaires isolés et la réutilisation de code.

**Cible** :
```
kafka-demo/
├── cmd/ (Points d'entrée)
│   ├── producer/main.go
│   ├── tracker/main.go
│   └── monitor/main.go
├── internal/ (Logique métier privée)
│   ├── kafka/ (Clients wrapper)
│   ├── processing/ (Logique de traitement)
│   └── monitor/ (Logique TUI)
├── pkg/ (Code réutilisable public)
│   └── models/
└── config/
```

### 1.2 Élimination des variables globales
**Priorité : Élevée**
Injecter les dépendances (Loggers, Config) via les constructeurs pour faciliter les tests et éviter les effets de bord.

---

## 2. ⚙️ Configuration et Environnement (Critique)

### 2.1 Fichier de configuration externe
**Priorité : Critique**
Remplacer les constantes hardcodées par un fichier `config.yaml` chargé au démarrage.
```yaml
app:
  env: "production"
kafka:
  broker: "kafka:9092"
  topic: "orders"
```

---

## 3. 🔄 Résilience et Fiabilité (Critique / Élevée)

### 6.1 Retry avec backoff exponentiel
**Priorité : Critique**
Le tracker doit pouvoir réessayer le traitement d'un message en cas d'erreur temporaire (ex: base de données inaccessible) avant d'abandonner.

### 6.3 Dead Letter Queue (DLQ)
**Priorité : Élevée**
Si un message échoue après X tentatives, il doit être envoyé vers un topic `orders-dlq` pour analyse manuelle, au lieu d'être perdu ou de bloquer la file.

### 6.2 Circuit Breaker
**Priorité : Moyenne**
Empêcher de surcharger un service en aval s'il est déjà en panne.

---

## 4. 🧪 Tests et Qualité (Élevée)

### 4.2 Amélioration de la couverture
**Priorité : Élevée**
Extraire la logique métier des fonctions `main()` vers des fonctions pures testables unitairement.

### 4.1 Tests d'intégration (Testcontainers)
**Priorité : Moyenne**
Utiliser Testcontainers pour lancer un vrai Kafka lors des tests `go test`, au lieu de mocker.

---

## 5. 🐳 Conteneurisation et Déploiement (Élevée)

### 7.1 Dockerfile multi-stage
**Priorité : Élevée**
Produire des images Docker légères (Alpine/Scratch) contenant uniquement le binaire compilé.

### 7.2 Docker Compose amélioré
**Priorité : Moyenne**
Ajouter Kafka-UI pour visualiser les messages facilement durant le développement.

---

## 6. 📊 Observabilité (Moyenne)

### 5.2 / 5.1 Prometheus & OpenTelemetry
**Priorité : Moyenne**
Le `log_monitor` TUI est excellent pour la démo, mais en production, l'export de métriques (endpoint `/metrics`) vers Prometheus/Grafana est le standard.

---

## 7. 🔒 Sécurité (Moyenne / Basse)

### 3.1 Authentification Kafka (SASL/SSL)
**Priorité : Moyenne**
Nécessaire si le cluster Kafka est partagé ou public.

---

## 8. 📝 Fonctionnalités Métier (Basse)

### 8.1 Support multi-topics
Extension pour gérer différents types d'événements.

### 8.2 Partitionnement intelligent
Utiliser le `customer_id` comme clé de partition pour garantir l'ordre des messages par client.

---

## 9. 🖥️ Interface & Divers (Basse)

### 9.1 / 9.2 Améliorations UI
Le moniteur actuel est suffisant pour le rôle de debugging. Une Web UI serait un projet à part entière.

### 12.1 Scripts PowerShell
Pour supporter les développeurs Windows nativement (actuellement WSL est recommandé).
