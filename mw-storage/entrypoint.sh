#!/bin/sh

go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -ext sql -dir internal/db/migration -seq init_schema

# Perform migrations
migrate -path internal/db/migration -database "postgresql://root:secret@postgres-storage:5432/mastersway_storage_db?sslmode=disable" -verbose up

# Start the server
./main
