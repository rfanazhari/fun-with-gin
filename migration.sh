#!/bin/bash

export DB_SCHEMA="apps"
export DB_USERNAME="fan"
export DB_PASSWORD="lupa_password"
export DB_NAME="gin_go"
export DB_HOST="localhost"
export DB_PORT=5432
export DB_TIMEZONE="Asia/Jakarta"
export DB_MAX_CON_IDLE=5
export DB_MAX_CON_OPEN=5
export DEBUG="true"
export APP_TIMEOUT=10
export NEED_SEEDER="true"
export MIGRATION_DIR="../../migrations/pgsql"

cd cmd/migration || exit
go run main.go
