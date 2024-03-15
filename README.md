# Fun with Gin-Gonic Golang

## Prerequisite

This project has the following prerequisites:

- [Golang v1.21.4](https://golang.org/doc/go1.21.4)
- [Gorm](https://gorm.io/) as ORM
- [PostgreSQL](https://www.postgresql.org/) as the database
- [JWT](https://jwt.io/) for token generation
- [Migrate](https://github.com/golang-migrate/migrate) as the engine for database migration
- [Gin Gonic](https://github.com/gin-gonic/gin) for REST HTTP

## Installation

1. sync go modules -> `go mod tidy`

#### Migration Database

1. download and install golang-migrate -> https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
2. migrate up service (e.g., service auth) -> \
   `migrate -source "file://migrations/pgsql" -database "postgres://username:pwd@localhost:5432/your_db?sslmode=disable&search_path=auth" up`
3. create new migration (if needed) (e.g., service auth) -> \
   `migrate create -ext sql -dir migrations/pgsql create_my_table`

## Run Apps

1. apps -> execute file `sh main.sh`
2. migration & seeder -> execute file `sh migration.sh`
