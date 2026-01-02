#!/bin/bash

# Test script for max-secure

set -e

echo "Running tests..."

# Run unit tests
echo "Running unit tests..."
go test -v -race -coverprofile=coverage.out ./...

# Show coverage
echo "Test coverage:"
go tool cover -func=coverage.out

# Generate HTML coverage report
echo "Generating coverage report..."
go tool cover -html=coverage.out -o coverage.html

echo "Tests complete! Coverage report: coverage.html"
