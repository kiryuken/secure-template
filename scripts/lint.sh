#!/bin/bash

# Linting and security scanning script

set -e

echo "Running linters and security scans..."

# Run golangci-lint
echo "Running golangci-lint..."
golangci-lint run ./...

# Run gosec for security scanning
echo "Running gosec..."
gosec ./...

# Run staticcheck
echo "Running staticcheck..."
staticcheck ./...

# Run govulncheck
echo "Running govulncheck..."
govulncheck ./...

echo "Linting and security scans complete!"
