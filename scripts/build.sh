#!/bin/bash

# Build script for max-secure

set -e

echo "Building max-secure..."

# Build API server
echo "Building API server..."
go build -o bin/api ./cmd/api

# Build worker
echo "Building worker..."
go build -o bin/worker ./cmd/worker

# Build CLI
echo "Building CLI..."
go build -o bin/cli ./cmd/cli

echo "Build complete! Binaries are in ./bin/"
