#!/bin/bash

export DB_SCHEMA=your_schema
export DB_USERNAME=your_username
export DB_PASSWORD=your_password
export DB_NAME=your_database
export DB_HOST=your_host
export DB_PORT=your_port
export DB_TIMEZONE=your_timezone
export DB_MAX_CON_IDLE=your_max_con_idle
export DB_MAX_CON_OPEN=your_max_con_open
export DEBUG=your_debug_value
export APP_TIMEOUT=your_app_timeout

cd cmd/migration || exit
go run main.go
