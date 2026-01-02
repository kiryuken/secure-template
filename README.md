# Max Secure Template

Enterprise-grade secure web application built with Go, featuring comprehensive security controls, audit logging, and modern DevOps practices.

## Features

### 🔐 Security

- **Authentication & Authorization**

  - JWT-based authentication
  - TOTP/MFA support
  - Role-based access control (RBAC)
  - Session management
  - Password reset & email verification

- **Data Protection**

  - AES-256-GCM encryption for data at rest
  - Argon2 password hashing
  - TLS/HTTPS support
  - Input validation & sanitization
  - XSS protection with HTML sanitization

- **Abuse Prevention**

  - Rate limiting (IP & user-based)
  - Request throttling
  - Circuit breaker pattern

- **Audit Trail**

  - Comprehensive audit logging
  - Before/after state tracking
  - User activity tracking
  - Export capabilities

### 🚀 Infrastructure

- **Database**

  - PostgreSQL with Bun ORM
  - Database migrations
  - Connection pooling

- **Caching**

  - Redis for session storage
  - Rate limit counters
  - Token blacklisting

- **File Storage**

  - AWS S3 integration
  - Local storage support
  - File type validation
  - Checksum verification

- **Background Jobs**

  - Async task processing with Asynq
  - Scheduled jobs with Cron
  - Email queue
  - File processing queue

### 📊 Observability

- **Logging**

  - Structured logging with Zap
  - Log rotation
  - Multiple log levels

- **Metrics**

  - Prometheus metrics
  - Request counters
  - Duration histograms
  - Custom business metrics

- **Tracing**

  - OpenTelemetry integration
  - Jaeger distributed tracing
  - Request tracing

- **Health Checks**

  - Liveness probes
  - Readiness probes
  - Dependency health checks

## Project Structure

```
max-secure/
├── cmd/                          # Main applications
│   ├── api/                      # API server
│   ├── worker/                   # Background worker
│   └── cli/                      # CLI tools
├── internal/                     # Private application code
│   ├── api/                      # API layer
│   │   ├── handlers/             # HTTP handlers
│   │   ├── middleware/           # HTTP middleware
│   │   └── routes/               # Route definitions
│   ├── domain/                   # Domain models
│   │   ├── models/               # Data models
│   │   └── entities/             # Business entities
│   ├── repository/               # Data access layer
│   ├── service/                  # Business logic
│   ├── auth/                     # Authentication (JWT, TOTP, passwords)
│   ├── crypto/                   # Encryption utilities
│   ├── validator/                # Input validation & sanitization
│   ├── storage/                  # File storage (S3, local)
│   ├── cache/                    # Redis cache
│   ├── queue/                    # Message queue
│   ├── email/                    # Email service
│   ├── worker/                   # Background job handlers
│   ├── audit/                    # Audit logging
│   └── notification/             # Notifications (Slack, etc.)
├── pkg/                          # Public reusable libraries
│   ├── logger/                   # Logging utilities
│   ├── errors/                   # Error handling
│   ├── config/                   # Configuration management
│   ├── database/                 # Database utilities
│   ├── telemetry/                # Tracing & metrics
│   ├── security/                 # Security utilities
│   ├── utils/                    # General utilities
│   ├── ratelimit/                # Rate limiting
│   ├── metrics/                  # Prometheus metrics
│   └── health/                   # Health checks
├── migrations/                   # Database migrations
├── configs/                      # Configuration files
├── deployments/                  # Deployment configs
│   ├── docker/                   # Docker configs
│   └── kubernetes/               # Kubernetes manifests
├── scripts/                      # Build & deployment scripts
├── tests/                        # Tests
│   ├── integration/              # Integration tests
│   ├── fixtures/                 # Test fixtures
│   └── mocks/                    # Test mocks
├── docs/                         # Documentation
│   ├── api/                      # API documentation
│   └── architecture/             # Architecture docs
├── .env.example                  # Environment variables example
├── docker-compose.yml            # Docker Compose config
├── Dockerfile                    # Docker image
├── Makefile                      # Build automation
├── go.mod                        # Go modules
└── README.md                     # This file
```

## Prerequisites

- Go 1.21+
- PostgreSQL 15+
- Redis 7+
- Docker & Docker Compose (for local development)

## Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/max-secure.git
cd max-secure
```

### 2. Setup environment variables

```bash
cp .env.example .env
# Edit .env with your configuration
```

### 3. Start dependencies with Docker Compose

```bash
docker-compose up -d postgres redis jaeger
```

### 4. Run database migrations

```bash
make migrate-up
# or
./scripts/migrate.sh up
```

### 5. Run the application

```bash
# Run API server
make run-api
# or
go run cmd/api/main.go

# Run worker (in separate terminal)
make run-worker
# or
go run cmd/worker/main.go
```

### 6. Access the application

- API: http://localhost:8080
- Health Check: http://localhost:8080/health
- Metrics: http://localhost:8080/metrics
- Jaeger UI: http://localhost:16686

## Development

### Build

```bash
make build
# or
./scripts/build.sh
```

### Run Tests

```bash
make test
# or
./scripts/test.sh
```

### Run Linters

```bash
make lint
# or
./scripts/lint.sh
```

### Database Migrations

```bash
# Create new migration
migrate create -ext sql -dir migrations -seq migration_name

# Run migrations
make migrate-up

# Rollback migration
make migrate-down
```

## Deployment

### Docker

```bash
# Build Docker image
docker build -t max-secure:latest .

# Run with Docker Compose
docker-compose up -d
```

### Kubernetes

```bash
# Apply Kubernetes manifests
kubectl apply -f deployments/kubernetes/

# Check deployment status
kubectl get pods
kubectl get services
```

## API Documentation

API documentation is available at `/swagger` when the server is running.

## Security Considerations

- Change all default passwords and secrets
- Use strong JWT secrets (32+ characters)
- Enable MFA for admin accounts
- Configure proper CORS origins
- Set appropriate rate limits
- Use HTTPS in production
- Regularly update dependencies
- Run security scans (`make lint`)
- Monitor audit logs

## Environment Variables

See [.env.example](.env.example) for all available environment variables.

Key variables:

- `DB_*`: Database configuration
- `REDIS_*`: Redis configuration
- `JWT_SECRET`: JWT signing secret
- `ENCRYPTION_KEY`: Data encryption key
- `AWS_*`: AWS/S3 configuration
- `SMTP_*`: Email configuration

## Testing

```bash
# Run all tests
make test

# Run with coverage
go test -cover ./...

# Run integration tests
go test -tags=integration ./tests/integration/...
```

## Monitoring

### Metrics (Prometheus)

Metrics are exposed at `/metrics`:

- HTTP request count
- Request duration
- Active connections
- Error rates
- Database query duration

### Tracing (Jaeger)

Distributed tracing is available through Jaeger UI at http://localhost:16686

### Logging

Logs are structured (JSON) and written to:

- Console (stdout)
- File (`./logs/app.log`)

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linters
5. Submit a pull request

## License

MIT License - see LICENSE file
