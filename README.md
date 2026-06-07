# 💎 Sims 4 Data Companion Platform

An enterprise-grade, microservice-based Data-as-a-Service (DaaS) platform designed to analyze, aggregate, and serve game mechanics metrics from *The Sims 4*. This project demonstrates high-performance backend patterns, rigorous security, distributed service architecture, and automated DevOps infrastructure.

---

## 🛠️ Tech Stack & Architecture

- **Core Backend Service:** Golang (Fiber v2) & gRPC Server
- **Travel Microservice:** Node.js (Express)
- **Database:** PostgreSQL (Relational Data Storage)
- **Caching Layer:** Redis (High-Performance Data Retrieval via Cache-Aside pattern)
- **Architecture:** Clean Architecture (SOLID Principles & Dependency Inversion)
- **Containerization:** Docker & Docker Compose
- **Data Engineering:** Python Automated ETL/DDL Seeder Script
- **Security:** JSON Web Token (JWT) Perimeter Authentication Middleware
- **Frontend:** HTML5, TailwindCSS (Glassmorphism UI), Vanilla JavaScript
- **API Documentation:** Swagger UI

---

## 🚀 Key Features

1. **Distributed Microservices (DaaS):** The main Go backend aggregates local database metrics with external data fetched seamlessly from a separate Node.js Travel API microservice.
2. **Dependency Inversion (SOLID):** Decoupled layers using Go Interfaces, ensuring the system remains highly testable and completely database-agnostic.
3. **High-Performance Caching:** Implemented Redis caching mechanism, providing blazing fast response times under 5ms.
4. **Advanced Global Error Handling:** Centralized custom middleware intercepts raw database errors and parses standardized JSON schemas securely.
5. **JWT Security Guard:** All telemetry data endpoints are restricted using cryptographically signed tokens.
6. **ETL Data Pipeline:** A Python-based seeder that extracts raw game data from CSVs, transforms, and loads it into the PostgreSQL database.
7. **Interactive API Docs:** Built-in Swagger UI for exploring and testing the API contract.

---

## 📦 How to Run the Infrastructure Locally

### 1. Clone & Spin up Containers
Ensure Docker is running on your machine.
```bash
git clone https://github.com/waranysye/sims-4-data-companion.git
cd sims-4-data-companion
docker-compose up -d --build
```
*This command will spin up the PostgreSQL database, Redis cache, Node.js Travel API, and the Go Backend.*

### 2. Run the Data Pipeline (ETL Seeder)
Run the Python script to inject the game data into the database.
```bash
# Ensure you have the psycopg2 module installed
pip install psycopg2
python seeder.py
```

### 3. Explore the Platform
- **Analytics Dashboard:** Open `index.html` in your web browser to view the premium dashboard.
- **Swagger Documentation:** Visit `http://localhost:8888/swagger` to explore the API.
- **Node.js Travel API:** Visit `http://localhost:3000/api/travel-packages`

---

## 🏛️ Project Structure

```text
.
├── config/              # Database & configuration setup
├── data/                # CSV files for the ETL pipeline
├── db/migrations/       # SQL migration scripts
├── handler/             # HTTP Handlers and gRPC servers
├── repository/          # Database interaction layer (Interfaces & Impl)
├── travel-api/          # Node.js Microservice for travel data
├── usecase/             # Business logic and data aggregation
├── main.go              # Go application entrypoint
├── seeder.py            # Python ETL script
├── docker-compose.yaml  # Docker services orchestration
├── index.html           # Frontend Analytics Dashboard
└── swagger.html         # Interactive API Documentation
```
