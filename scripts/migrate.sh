#!/bin/bash

# Migration script

set -e

COMMAND=${1:-up}

echo "Running migrations: $COMMAND"

# Run migrations using golang-migrate
migrate -path ./migrations -database "postgresql://postgres:postgres@localhost:5432/max_secure?sslmode=disable" $COMMAND

echo "Migrations complete!"
