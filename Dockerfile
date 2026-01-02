# Multi-stage build for production
FROM golang:1.21-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git make ca-certificates tzdata

WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api ./cmd/api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o worker ./cmd/worker
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o cli ./cmd/cli

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy binaries from builder
COPY --from=builder /app/api .
COPY --from=builder /app/worker .
COPY --from=builder /app/cli .

# Copy migrations
COPY --from=builder /app/migrations ./migrations

# Copy configs
COPY --from=builder /app/configs ./configs

# Expose port
EXPOSE 8080

# Run the application
CMD ["./api"]
