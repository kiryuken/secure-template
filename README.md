# Max Secure Template

🚀 **Production-ready Go boilerplate template** for building enterprise-grade secure web applications with comprehensive security controls, audit logging, and modern DevOps practices.

> **⚠️ This is a template with TODO placeholders** - You need to implement the business logic yourself. All files contain structured comments and implementation guides to help you develop your application quickly.

## 📋 What's Included

This template provides a complete project structure with:

- ✅ **Well-organized folder structure** following clean architecture
- ✅ **Dependency injection patterns** with clear interfaces
- ✅ **Detailed TODO comments** explaining what to implement
- ✅ **Security best practices** (JWT, MFA, encryption, rate limiting)
- ✅ **Production-ready configs** (Docker, Kubernetes, CI/CD)
- ✅ **Modular code** with max 80 lines per file for maintainability

## 🎯 Quick Start

1. **Clone this template**
2. **Search for `// TODO:` comments** across the codebase
3. **Implement your business logic** following the provided structure
4. **Configure environment variables** from `.env.example`
5. **Run and test** your application

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
git clone https://github.com/kiryuken/secure-template.git
cd secure-template
```

### 2. Understand the structure

All implementation files contain `// TODO:` comments with detailed steps:

```go
// Example from handlers/auth_login.go
func (h *AuthHandler) Login(ctx context.Context) error {
    // TODO: Implement
    // - Validate email & password from request
    // - Call authService.Login()
    // - Return JWT tokens
    // - Log audit trail
    return nil
}
```

### 3. Implement your business logic

Search for `// TODO:` in your IDE to find all implementation points:

- **Handlers**: HTTP request handling
- **Services**: Business logic layer
- **Repositories**: Database queries with Bun ORM
- **Middleware**: Authentication, rate limiting, security headers

### 4. Setup environment variables

```bash
cp .env.example .env
# Edit .env with your configuration
```

### 5. Start dependencies with Docker Compose

```bash
docker-compose up -d postgres redis jaeger
```

### 6. Install Go dependencies

```bash
go mod download
go mod tidy
```

### 7. Implement required functionality

Before running, you need to implement:

- Database connection in `pkg/database/`
- Configuration loading in `pkg/config/`
- Main server setup in `cmd/api/main.go`
- At minimum: one working handler + service + repository

### 8. Run database migrations

```bash
# After implementing migration logic
make migrate-up
```

### 9. Run the application

```bash
# After implementing main.go and core dependencies
go run cmd/api/main.go
```

## 📂 Implementation Guide

### Key Files to Implement

1. **Configuration** (`pkg/config/config.go`)

   - Load environment variables
   - Validate configuration
   - Set defaults

2. **Database** (`pkg/database/postgres.go`)

   - Setup Bun ORM connection
   - Connection pooling
   - Health checks

3. **Main Server** (`cmd/api/main.go`)

   - Initialize dependencies
   - Setup HTTP server
   - Register routes and middleware
   - Graceful shutdown

4. **Handlers** (`internal/api/handlers/`)

   - Implement HTTP request handling
   - Input validation
   - Response formatting

5. **Services** (`internal/service/`)

   - Implement business logic
   - Call repositories
   - Handle transactions

6. **Repositories** (`internal/repository/`)
   - Implement database queries with Bun ORM
   - CRUD operations
   - Complex queries with filters

### Example Implementation Flow

1. **Start with User Management:**

   ```
   models/user_model.go →
   repository/user_repository.go →
   service/user_service.go →
   handlers/user_*.go
   ```

2. **Add Authentication:**

   ```
   auth/jwt.go →
   auth/password.go →
   service/auth_service.go →
   handlers/auth_*.go →
   middleware/auth.go
   ```

3. **Add Supporting Features:**
   - Rate limiting (middleware/ratelimit.go)
   - Audit logging (audit/audit.go)
   - File uploads (storage/s3.go)

## 📖 Code Structure

All template files follow this pattern:

```go
type Service struct {
    // TODO: Add dependencies
    // repo   repository.Repository
    // logger logger.Logger
}

func NewService( /* dependencies */ ) *Service {
    return &Service{
        // TODO: Initialize dependencies
    }
}

func (s *Service) Method( /* parameters */ ) error {
    // TODO: Implement
    // - Step 1: Validate input
    // - Step 2: Process data
    // - Step 3: Save to database
    // - Step 4: Return response
    return nil
}
```

## 🔍 Finding TODOs

Use your IDE to search for TODO comments:

- **VS Code**: `Ctrl+Shift+F` → search for `// TODO:`
- **GoLand**: `View → Tool Windows → TODO`
- **Command line**: `grep -r "// TODO:" internal/ pkg/`

## 📦 Dependencies to Install

After implementing core logic, add these packages:

```bash
go get github.com/uptrace/bun
go get github.com/uptrace/bun/driver/pgdriver
go get github.com/golang-jwt/jwt/v5
go get github.com/pquerna/otp
go get golang.org/x/crypto/argon2
go get github.com/go-redis/redis/v8
go get github.com/aws/aws-sdk-go-v2
go get go.uber.org/zap
go get github.com/hibiken/asynq
# ... see go.mod for complete list
```

## 🎨 Customization

This template is designed to be customized:

- **Add your models** in `internal/domain/models/`
- **Add your DTOs** in `internal/api/dto/`
- **Add your handlers** in `internal/api/handlers/`
- **Add your services** in `internal/service/`
- **Add your repositories** in `internal/repository/`
- **Keep files under 80 lines** for maintainability

## ⚙️ Configuration

After implementing the application, it will be available at:

- API: http://localhost:8080
- Health Check: http://localhost:8080/health
- Metrics: http://localhost:8080/metrics
- Jaeger UI: http://localhost:16686

## 🛠️ Development Commands

### Build (after implementation)

```bash
make build
```

### Run Tests (after writing tests)

```bash
make test
```

### Run Linters

```bash
make lint
```

### Database Migrations (after implementation)

```bash
# Create new migration
migrate create -ext sql -dir migrations -seq migration_name

# Run migrations
make migrate-up

# Rollback migration
make migrate-down
```

## 🚀 Deployment (Production Ready)

This template includes production-ready deployment configurations.

### Docker

```bash
# Build Docker image
docker build -t secure-app:latest .

# Run with Docker Compose
docker-compose up -d
```

### Kubernetes

```bash
# Apply Kubernetes manifests
kubectl apply -f deployments/kubernetes/

# Check deployment status
kubectl get pods
```

## 📚 Documentation

- **[STRUCTURE.md](docs/STRUCTURE.md)**: Detailed project structure explanation
- **API Docs**: Available at `/swagger` endpoint (after implementation)
- **Architecture**: See `docs/architecture/` directory

## ✅ Checklist for Implementation

Use this checklist to track your progress:

- [ ] Implement configuration loading (`pkg/config/`)
- [ ] Setup database connection (`pkg/database/`)
- [ ] Implement logger setup (`pkg/logger/`)
- [ ] Create your domain models (`internal/domain/models/`)
- [ ] Implement repositories with Bun ORM (`internal/repository/`)
- [ ] Implement service layer (`internal/service/`)
- [ ] Implement HTTP handlers (`internal/api/handlers/`)
- [ ] Setup middleware (`internal/api/middleware/`)
- [ ] Configure routes (`internal/api/routes/`)
- [ ] Implement JWT authentication (`internal/auth/`)
- [ ] Setup Redis caching (`internal/cache/`)
- [ ] Implement audit logging (`internal/audit/`)
- [ ] Add rate limiting
- [ ] Write unit tests
- [ ] Write integration tests
- [ ] Setup CI/CD pipeline
- [ ] Configure production environment
- [ ] Review security settings

## 🔒 Security Reminders

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
- `JWT_SECRET`: JWT signing secret (32+ characters)
- `ENCRYPTION_KEY`: Data encryption key (32 bytes)
- `AWS_*`: AWS/S3 configuration
- `SMTP_*`: Email configuration

Example `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your-password
DB_NAME=secure_db

JWT_SECRET=your-strong-secret-key-min-32-chars
JWT_EXPIRY=15m

REDIS_HOST=localhost
REDIS_PORT=6379

ENCRYPTION_KEY=your-32-byte-encryption-key
```

## 🧪 Testing (After Implementation)

```bash
# Run all tests
make test

# Run with coverage
go test -cover ./...

# Run integration tests
go test -tags=integration ./tests/integration/...
```

## 📊 Monitoring (After Implementation)

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

## 💡 Tips for Development

1. **Start Small**: Implement one feature at a time (e.g., user CRUD first)
2. **Follow the TODOs**: Each file has step-by-step implementation guides
3. **Test as You Go**: Write tests for each component you implement
4. **Use the Structure**: Don't fight the architecture, follow the patterns
5. **Keep Files Small**: Stick to the 80-line guideline for maintainability
6. **Check Examples**: See `docs/STRUCTURE.md` for more detailed explanations

## 🤝 Contributing

This is a template repository. To contribute:

1. Fork the repository
2. Create a feature branch
3. Improve the template structure or add documentation
4. Run tests and linters
5. Submit a pull request

## 📝 License

MIT License - see LICENSE file

## 🔗 Links

- **Repository**: https://github.com/kiryuken/secure-template
- **Issues**: Report bugs or request features
- **Discussions**: Ask questions or share your implementations

---

**Remember**: This is a template with TODO placeholders. You need to implement the actual business logic. Happy coding! 🚀
