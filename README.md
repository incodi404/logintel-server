# Logintel Central Server

Logintel is an open source security monitoring and control system. This is used for kernel-level monitoring, command-and-control, alert generating based on rules, storing and visualizing logs efficiently.

Logintel Central Server is a SIEM System who is responsible for collecting logs from [Logintel Agent](https://github.com/incodi404/logintel), storing the logs efficiently in Elasticsearch with 3 hours of TTL, connecting with Kibana, generating alerts based on rules, admin and agent management and C2 system.

#### The system is still under development and getting better day by day. The first release of the system will be within October, 2026.

## Why A Central Server?

[Logintel Agent](https://github.com/incodi404/logintel) collects events and that's not enough. The logs should be processed and stored efficiently. If anything goes wrong, instead of roaming file to file to get logs, required logs should be at one place. Storing is also not enough, there should be a dynamic rule engine which checks every log and alert before everything goes wrong. It is not only a SIEM system but also integrated command-and-control system which allows the admin to access endpoint system's shell without any extra authentication or any type of fingerprints. The connection is instant and secured with mTLS and also there will be no constant open port.

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

## Technologies

### Golang

The entire project, agent and server, are built with Golang. Golang is very lightweight and well-connected to Linux kernel system and has well-maintained packages of eBPF that makes it the best choice for the agent. Also, Golang has very good performance in backend that most of the big companies are now using it in their backend and building cloud native tools with it. After considering all these points, it was crystal clear that Golang is the best choice for the system.

### eBPF

eBPF is a technology that let developers run custom code within the kernel without writing any kernel module. It also provide enough security to handle bugs in the kernel, so the kernel will not crash. In this project, eBPF is used to capture kernel events that creates a transparancy which is the best way to understand what actually happend and how it was happened.

### fanotify

Fanotify is a kernel notification subsystem that notifies user-space applications when any file events occure in the kernel. It keeps eye on the entire filesystem. It is one of the best choice to track file operarions in a system.

#### Trade-off

It produces massive amount of logs. A good amount of logs have no connection with security events.

#### Implemented Solution

A file path-based blacklist filtration is implemented in the agent. A log with blacklisted path will never reach gRPC. The list is configurable with a YAML file.

### Dbus IPC

Dbus Inter-Process Communication is a system in which user-applications get notified about status of systemd services. Well, this is not connected with any security purpose but we integrated it to get notified whenever anything goes wrong with any systemd service including Logintel Agent.

### Elasticsearch

Elasticsearch is used to store logs and suspicious logs. With low-latency filtration, in-built rollover system, auto-deletion with ILM make the database perfect fit for the system.

### Kibana

Kibana provides stored log visualization with table, graphs and filtration without any extra-coding. One of the best choice for log visualization with in-built integration Elasticsearch REST APIs.

### PostgreSQL

What else should we choose for RDBMS. RBAC structure, command history, agent management all are handled by this database. Chose for scalability and robustness.

### NATS Jetstream

NATS JetStream is the persistence and streaming layer built into NATS, a high-performance messaging system used in distributed applications and microservices. While Core NATS is an in-memory pub/sub system that delivers messages only to active subscribers, JetStream adds durable storage, message replay, acknowledgments, and delivery guarantees.

Integration of Apache Kafka is heavy, latency of RabbitMQ pub/sub is painful. NATS Jetstream provides speed, parsistant storage, guarantee and that are the reason it is here.

### RabbitMQ

RabbitMQ handles the background jobs here. For now, it is used to send email on every alert.

### React JS

React JS is for admin panel.

## Future integration

There is a planning to add a machine learning-based correlation engine and a RAG system to analysis suspicious logs.

## Kibana Dashboard & Log Visualization

### Unified Log Table

![Unified Log Table](https://res.cloudinary.com/fwkfpmra/image/upload/v1785757989/Screenshot_2026-07-18_154200_xvaxjq.png)

### Single Resource Log

![Single Resource Log](https://res.cloudinary.com/fwkfpmra/image/upload/v1785757857/Screenshot_2026-07-18_153906_b8pds9.png)

### Single Log Details

![Single Log Details](https://res.cloudinary.com/fwkfpmra/image/upload/v1785757990/Screenshot_2026-07-18_154314_nkwr0e.png)

### Agent Running in Background

![Agent Running in Background](https://res.cloudinary.com/fwkfpmra/image/upload/v1785757989/Screenshot_2026-07-18_174945_fqpfb2.png)

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
