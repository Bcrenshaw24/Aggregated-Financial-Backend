## Distributed Financial Analytics 

A containerized financial backend for aggregating account and transaction data, tracking balances, and analyzing spending patterns. Built with Go, PostgreSQL, Docker, and gRPC.

> **Status:** In development

## Overview

The platform integrates with external financial data providers to retrieve account and transaction information. It persists financial records in PostgreSQL and provides services for transaction management, spending analytics, and anomaly detection.

### Key Features

- **Account management:** Track linked financial accounts and their reported balances.
- **Transaction ingestion:** Synchronize, deduplicate, and persist external transaction data.
- **Spending analytics:** Aggregate spending over weekly and monthly time windows.
- **Anomaly detection:** Identify potentially unusual spending patterns.
- **Service communication:** Use gRPC for internal communication between backend services.
- **Observability:** Integrate distributed tracing, metrics, and structured logging.
- **Containerization:** Run application services and infrastructure with Docker Compose.

*Features listed here represent project goals; implementation status will evolve.*

## Architecture

The platform is organized into independently running services:

| Service | Responsibility |
|---|---|
| API Gateway | Entry point for client requests and routing to backend services |
| Account Service | Manage linked accounts and account balances |
| Transaction Service | Synchronize, deduplicate, and retrieve transactions |
| Analytics Service | Calculate spending summaries and detect anomalies |
| PostgreSQL | Persist accounts, transactions, and synchronization state |

External financial providers are accessed through dedicated integration clients. Internal services communicate using gRPC.

## Repository Structure

```text
finance-backend/
├── services/
│   ├── api-gateway/
│   ├── account-service/
│   ├── transaction-service/
│   └── analytics-service/
├── shared/
│   ├── proto/
│   └── go/
├── db/
│   └── migrations/
├── infra/
│   ├── postgres/
│   ├── otel-collector/
│   ├── prometheus/
│   └── grafana/
└── docker-compose.yml
```

### Directory Guide

- `services/` — Independently deployable Go applications, each with its own entry point and service-specific implementation.
- `services/api-gateway/` — Handles incoming client requests and forwards them to the appropriate backend service.
- `services/account-service/` — Manages account metadata and reported balances.
- `services/transaction-service/` — Integrates with external financial APIs and manages transaction synchronization and persistence.
- `services/analytics-service/` — Performs spending aggregations and anomaly analysis.
- `shared/proto/` — Shared Protocol Buffer definitions for internal gRPC APIs.
- `shared/go/` — Reusable Go packages shared across services.
- `db/migrations/` — Versioned SQL migrations for creating and evolving the database schema.
- `infra/postgres/` — PostgreSQL-specific configuration and initialization resources.
- `infra/otel-collector/` — OpenTelemetry Collector configuration.
- `infra/prometheus/` — Prometheus metrics collection configuration.
- `infra/grafana/` — Grafana dashboards and provisioning configuration.
- `docker-compose.yml` — Defines the local multi-container environment.

## Technology Stack

| Technology | Purpose |
|---|---|
| Go | Backend services |
| PostgreSQL | Relational persistence and analytical queries |
| gRPC / Protocol Buffers | Internal service communication |
| Docker / Docker Compose | Containerization and local orchestration |
| OpenTelemetry | Instrumentation and distributed tracing |
| Prometheus / Grafana | Metrics and visualization |

## Getting Started

### Prerequisites

- Go
- Docker Engine or Docker Desktop
- Docker Compose

### Run the development environment

```bash
docker compose up --build -d
```

Check container status:

```bash
docker compose ps
```

View service logs:

```bash
docker compose logs -f
```

Stop the environment:

```bash
docker compose down
```

> Setup instructions are preliminary. Database migrations, environment variables, and service startup requirements must be configured before the full environment can run successfully.

## Roadmap

- [ ] Establish the Docker Compose development environment.
- [ ] Implement database migrations and connection pooling.
- [ ] Define gRPC service contracts.
- [ ] Integrate the external financial data provider.
- [ ] Implement idempotent transaction synchronization.
- [ ] Add spending analytics and anomaly detection.
- [ ] Instrument services with OpenTelemetry.
- [ ] Add integration tests and database performance benchmarks.

## License

License to be determined.
