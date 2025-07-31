# 🎉 Fun with Gin-Gonic Golang

A modern Golang project template leveraging the **Gin-Gonic** framework, built for rapid RESTful API development with robust tooling and best practices.

## 🚀 Overview

Kickstart your Go backend journey with a full-stack ready base offering:

- Clean REST API architecture with **Gin Gonic**
- Reliable database management using **GORM** & **PostgreSQL**
- Secure authentication with **JWT**
- Zero-downtime migrations via **golang-migrate**
- Fast session/token handling with **Redis**
- Effortless local or Docker-based deployments

## 📦 Prerequisites

Before getting started, ensure you have:

- [Golang v1.21.4](https://golang.org/doc/go1.21.4)
- [GORM](https://gorm.io/) (ORM)
- [PostgreSQL](https://www.postgresql.org/) (database)
- [JWT](https://jwt.io/) (authentication tokens)
- [golang-migrate](https://github.com/golang-migrate/migrate) (DB migrations)
- [Gin Gonic](https://github.com/gin-gonic/gin) (HTTP framework)
- [Redis](https://github.com/redis/go-redis) (session storage)

## ⚡ Quick Start

### 1. Sync Go Modules

```bash
go mod tidy
```

## 🗄️ Database Migration

**Install migration tool:**

- Download [golang-migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

**Migrate schema:**

```bash
migrate -source "file://migrations/pgsql" -database "postgres://username:pwd@localhost:5432/your_db?sslmode=disable&search_path=apps" up
```

**Create a new migration:**

```bash
migrate create -ext sql -dir migrations/pgsql create_my_table
```

## 🏃 Running the App

**1. Run Application:**
```bash
sh main.sh
```

**2. Migration & Seeder:**
```bash
sh migration.sh
```

**3. Using Docker Compose:**
```bash
docker-compose up -d
```

**4. Manual migration in Docker:**
```bash
docker ps
docker exec -it  bash
cd cmd/migration
go run main.go
```

> Modern, scalable, and developer-friendly: unleash the power of Gin-Gonic and Go for your next project! Fork, star, and give feedback on GitHub.

[1] https://golang.org/doc/go1.21.4
[2] https://gorm.io
[3] https://www.postgresql.org
