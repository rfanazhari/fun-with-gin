#!/bin/bash

# Set environment variables
export DB_SCHEMA=apps
export DB_USERNAME=fan
export DB_PASSWORD=lupa_password
export DB_NAME=gin_go
export DB_HOST=localhost
export DB_PORT=4567
export DB_TIMEZONE=Asia/Jakarta
export DB_MAX_CON_IDLE=10
export DB_MAX_CON_OPEN=10
export DEBUG=false
export APP_TIMEOUT=10
export APP_PORT=:3001
export PROXY_URL="127.0.0.1;::1"

# Run the Go program
go run cmd/main.go
