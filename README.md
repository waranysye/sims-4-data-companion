\# 💎 Sims 4 Data Companion Engine



An Enterprise-Grade Full-Stack Data-as-a-Service (DaaS) platform built to analyze and extract game mechanics metrics from The Sims 4. This project demonstrates high-performance backend patterns, rigorous security parameters, and automated DevOps infrastructure.



\---



\## 🛠️ Tech Stack \& Architecture

\- \*\*Backend:\*\* Golang (Go Fiber v2) \& gRPC Server (`:50051`)

\- \*\*Database:\*\* PostgreSQL (Relational Data Storage)

\- \*\*Caching Layer:\*\* Redis Cache (High-Performance Data Retrieval)

\- \*\*Architecture:\*\* Clean Architecture (SOLID Principles \& Dependency Inversion)

\- \*\*Containerization:\*\* Docker \& Docker Compose

\- \*\*Data Ingestion:\*\* Python Automated ETL/DDL Seeder Script

\- \*\*CI/CD Pipeline:\*\* GitHub Actions (Automated Integration Testing)

\- \*\*Security:\*\* JSON Web Token (JWT) Perimeter Authentication Middleware



\---



\## 🚀 Key Features Demonstrated

1\. \*\*Dependency Inversion (SOLID):\*\* Decoupled layers using Go Interfaces, ensuring the system remains completely database-agnostic.

2\. \*\*High-Performance Caching:\*\* Implemented Redis caching mechanism with Cache-Aside pattern, blazing fast response time under 5ms.

3\. \*\*Advanced Global Error Handling:\*\* Centralized custom middleware intercepts raw database errors and parses standardized JSON schemas securely.

4\. \*\*JWT Security Guard:\*\* All telemetry data endpoints are restricted using cryptographically signed tokens.

5\. \*\*Automated CI DevOps:\*\* GitHub Actions runner provisions cloud environments to spin up tests automatically on every code commit.



\---



\## 📦 How to Run the Infrastructure Locally



\### 1. Clone \& Spin up Containers

```bash

git clone \[https://github.com/waranysye/sims-4-data-companion.git](https://github.com/waranysye/sims-4-data-companion.git)

cd sims-4-data-companion

docker-compose up -d

