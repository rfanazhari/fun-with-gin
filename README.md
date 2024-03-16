# Fun with Gin-Gonic Golang

## Prerequisite

This project has the following prerequisites:

- [Golang v1.21.4](https://golang.org/doc/go1.21.4)
- [Gorm](https://gorm.io/) as ORM
- [PostgreSQL](https://www.postgresql.org/) as the database
- [JWT](https://jwt.io/) for token generation
- [Migrate](https://github.com/golang-migrate/migrate) as the engine for database migration
- [Gin Gonic](https://github.com/gin-gonic/gin) for REST HTTP
- [Redis](https://github.com/redis/go-redis) for storing session login tokens

## Installation

1. sync go modules 
   ```bash 
   go mod tidy 
   ```

#### Migration Database

1. download and install golang-migrate -> https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
2. migrate up
   ```bash
   migrate -source "file://migrations/pgsql" -database "postgres://username:pwd@localhost:5432/your_db?sslmode=disable&search_path=apps" up 
   ```
3. create new migration (if needed)
   ```bash
   migrate create -ext sql -dir migrations/pgsql create_my_table
   ```

## Run Apps

1. apps -> execute file 
   ```bash
   sh main.sh
   ```
2. migration & seeder -> execute file
   ```bash
   sh migration.sh
   ```
3. using docker-compose
   ```bash
   docker-compose up -d
   ```
4. if using docker the migration should execute inside the container
   ```bash
   docker ps
   docker exec -it CONTAINER-ID bash
   cd cmd/migration
   go run main.go
   ```
