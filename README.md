# Logintel Central Server

Logintel is an open source security monitoring and control system. This is used for kernel-level monitoring, command-and-control, alert generating based on rules, storing and visualizing logs efficiently.

Logintel Central Server is a SIEM System who is responsible for collecting logs from [Logintel Agent](https://github.com/incodi404/logintel), storing the logs efficiently in Elasticsearch with 3 hours of TTL, connecting with Kibana, generating alerts based on rules, admin and agent management and C2 system.

#### The system is still under development and getting better day by day. The first release of the system will be within October, 2026.

## Features

- Ingest logs from different agents
- Store logs in Elasticsearch for 3 hours
- Log visualization in Kibana
- Stream logs to NATS Jetstream
- Rule Engine collects logs from NATS Jetstream, apply rules and raise alerts if get anything suspicious
- Store alert history in PostgreSQL and Elasticsearch for persistant storage
- Admin panel to create new agents, view alert history etc.
- C2 system

## Services

### Log Ingestion Service

Log ingestion service is responsible for collecting logs from different services via **gRPC**. After collecting logs, it saves the logs in Elasticsearch in different data streams and stream the logs in NATS Jetstream Pub/Sub for fan-out distribution.

### Kibana

Kibana dashboard is connected with Elasticseach and use for log visualization, filtration and exporting as CSV.

### Rule Engine

Rule engine is the heart of this system. This will load rules from PGSQL in every 5 minutes to its memory. It check every log against the rules. If anything suspicious detects then the log will saved with rule metadata in Elasticsearch and PGSQL. Then it will go to RabbitMQ for email alert.

### Auth & Agent Management

This is the main service that is used for admin authentication and authorization and agent management. Adding new agent, updating data of an agent, everything will be done here. It will have REST APIs.

### C2 System

The agent will have a C2 system that allows the admin to run command directly from browser to shell without any SSH fingerprint and extra authentication. The connection will be secured by mTLS. It is the incident response functionality that will be integrated with the agent.

## Architecture

![Architecture](https://res.cloudinary.com/fwkfpmra/image/upload/v1785779179/server.arch_acqvpa.png)

## Roadmap for v1

- ✅ Log ingestion from different agents
- ✅ Storing logs in Elasticsearch
- ✅ Visualizing logs in Kibana
- ⬜ Rule Engine
- ⬜ C2 System
- ⬜ Admin Panel
- ⬜ Agent-Server Authentication

## Spin-Up Server with Docker

```shell
git clone https://github.com/incodi404/logintel-server.git
cd logintel-server
docker compose up
```

## Author

Dipankar Chowdhury — Backend Engineer (Security Focused) | [GitHub](https://github.com/incodi404/) · [LinkedIn](https://www.linkedin.com/in/dipankar-chowdhury/)
