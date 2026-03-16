# Go-Monitor
 
Distributed monitoring system in Go.

## Features

- CPU / RAM / Disk / Load / Network metrics
- Distributed agents
- REST API + WebSocket live metrics
- Simple web dashboard
- Alerts (console, Telegram, Slack)
- Docker & docker-compose support
- PostgreSQL or SQLite storage

## Architecture

Agents → Message Queue → Server → Database → Web Dashboard

## 📂 Project structure
```bash
go-monitor/
├── agent/
│   ├── main.go
│   ├── metrics.go
│   ├── sender.go
│   └── config.yaml
├── server/
│   ├── main.go
│   ├── api.go
│   ├── alerts.go
│   ├── storage.go
│   ├── websocket.go
│   └── models.go
├── web/
│   └── dashboard.html
├── config/
│   └── server.yaml
├── docker/
│   ├── Dockerfile.agent
│   ├── Dockerfile.server
│   └── docker-compose.yaml
└── README.md
```
## Getting Started

### Run with Docker

```bash
docker-compose up --build
```
## Run Locally

1.Start PostgreSQL

2.Start server:

```bash
cd server
go run main.go
```
3.Start agent:

```bash
cd agent
go run main.go
```

4. Open dashboard: http://localhost:8080

## Alerts

- CPU > 90%

- Memory > 85%

- Disk > 90%

# 🔹 How to use on GitHub

1. Create a `go-monitor` repository.
2. Copy all files, including the `agent`, `server`, `web`, `config`, and `docker` folders.
3. Add `.gitignore` for Go and Docker.
4. Push to GitHub:
```bash
git init
git add .
git commit -m "Initial commit Go-Monitor"
git branch -M main
git remote add origin https://github.com/mscbuild/Go-Monitor.git
git push -u origin main
