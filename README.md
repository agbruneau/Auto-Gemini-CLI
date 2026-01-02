# 🚀 Kafka Order Tracking System

[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Kafka](https://img.shields.io/badge/Apache_Kafka-3.7.0-white?style=flat&logo=apache-kafka)](https://kafka.apache.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A robust, enterprise-grade **Event-Driven Architecture (EDA)** demonstration using **Go** and **Apache Kafka**. This project simulates a complete e-commerce order lifecycle—from generation to real-time tracking—featuring high observability via a dedicated Terminal User Interface (TUI).

---

## 🏗 System Architecture

The ecosystem consists of three decoupled core services:

1.  **📦 Producer (`producer`)**: Simulates customer activity by generating enriched order events and streaming them to the `orders` Kafka topic.
2.  **⚙️ Tracker (`tracker`)**: Consumes order events in real-time, performing validation and maintaining a comprehensive audit trail.
3.  **📊 Monitor (`log_monitor`)**: A sophisticated TUI dashboard providing live visualization of system performance, throughput, and success rates.

---

## 🌟 Key Features & Design Patterns

This implementation adheres to modern distributed systems standards:

- **Event-Driven Architecture (EDA)**: Complete decoupling of services through asynchronous messaging.
- **Event Carried State Transfer (ECST)**: Self-contained messages that include all necessary context (product, customer, pricing), minimizing downstream lookups.
- **Guaranteed Delivery**: Implements Kafka delivery reports (ACKs) to ensure data integrity.
- **Dual-Stream Observability**:
  - **Technical Health**: Structured JSON logging (`tracker.log`) for system monitoring.
  - **Business Audit**: An immutable event journal (`tracker.events`) for compliance and debugging.
- **Graceful Shutdown**: Strict handling of `SIGTERM`/`SIGINT` signals for zero-data-loss termination.
- **Operational Idempotence**: Automated infrastructure setup via robust shell orchestration.

---

## 🛠 Prerequisites

Ensure the following are installed on your system:

1.  **Docker** and **Docker Compose** (V2).
2.  **Go** (version 1.22 or higher).
3.  An **ANSI-compatible terminal** (for the TUI monitor).
4.  **Sudo privileges** (required for Docker commands in scripts).

---

## 🚀 Getting Started

Deploy the complete environment with a single command:

```bash
./start.sh
```

**Automated actions performed:**

- Deploys a Kafka cluster (KRaft mode) via Docker.
- Polls the broker until health checks pass.
- Idempotently creates the `orders` topic.
- Launches the **Producer** and **Tracker** services in the background.

---

## 📊 Monitoring & Observation

### 1. Interactive Dashboard (Recommended)

Launch the TUI monitor in a **new terminal window** for real-time visualization:

```bash
go run -tags monitor cmd_monitor.go log_monitor.go models.go constants.go
```

- **Controls**: Press `q` or `Ctrl+C` to exit.
- **Insights**: Monitor msg/sec throughput, success rates, and live log streams.

### 2. Manual Log Inspection

Follow the generated logs directly:

```bash
# Business Audit Trail
tail -f tracker.events

# Technical System Logs (Formatted with jq)
tail -f tracker.log | jq
```

---

## 🛑 Stopping the System

To gracefully terminate all services and the Kafka infrastructure:

```bash
./stop.sh
```

_This script uses PID tracking to send `SIGTERM` signals, allowing Go services to flush buffers and close connections properly before the Docker environment is torn down._

---

## 📂 Project Structure

- **Entry Points (`cmd_*.go`)**:
  - `cmd_producer.go`: Producer initialization.
  - `cmd_tracker.go`: Tracker initialization.
  - `cmd_monitor.go`: TUI Monitor initialization.
- **Core Logic**:
  - `producer.go`: Kafka publishing and order generation.
  - `tracker.go`: Message consumption and processing.
  - `log_monitor.go`: TUI widgets and metrics logic.
- **Shared Resources**:
  - `models.go`: Structured log and event definitions.
  - `order.go`: The `Order` domain model (ECST).
  - `constants.go`: Global configuration (Topics, Timeouts).
- **Orchestration**:
  - `start.sh` / `stop.sh`: Lifecycle management scripts.
  - `docker-compose.yaml`: Infrastructure definition.

---

## 💻 Development & Testing

The project uses Go **Build Tags** (`producer`, `tracker`, `monitor`, `kafka`) for modular compilation.

### Manual Compilation

```bash
# Build Producer
go build -tags producer -o producer cmd_producer.go producer.go order.go models.go constants.go

# Build Tracker
go build -tags tracker -o tracker cmd_tracker.go tracker.go order.go models.go constants.go

# Build Monitor
go build -tags monitor -o monitor cmd_monitor.go log_monitor.go models.go constants.go
```

### Running Tests

```bash
# Producer Tests
go test -tags kafka,producer producer.go producer_test.go order.go constants.go

# Tracker Tests
go test -tags kafka,tracker tracker.go tracker_test.go order.go constants.go models.go

# Monitor Tests
go test -tags monitor log_monitor.go log_monitor_test.go models.go constants.go
```
