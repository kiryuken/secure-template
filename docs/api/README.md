# API Documentation

## Authentication Endpoints

### POST /auth/login

Login with credentials

**Request:**

```json
{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**

```json
{
  "access_token": "eyJhbGc...",
  "refresh_token": "eyJhbGc...",
  "expires_in": 900
}
```

### POST /auth/register

Register new user

**Request:**

```json
{
  "email": "user@example.com",
  "username": "username",
  "password": "password123",
  "first_name": "John",
  "last_name": "Doe"
}
```

### POST /auth/logout

Logout and invalidate session

### POST /auth/refresh

Refresh access token

### POST /auth/forgot-password

Request password reset

### POST /auth/reset-password

Reset password with token

### POST /auth/enable-mfa

Enable MFA/TOTP

### POST /auth/verify-mfa

Verify TOTP code

## User Endpoints

### GET /users/:id

Get user by ID

### GET /users

List users (with pagination)

### PUT /users/:id

Update user

### DELETE /users/:id

Delete user

### GET /profile

Get current user profile

### PUT /profile

Update current user profile

### POST /profile/avatar

Upload user avatar

## File Endpoints

### POST /files

Upload file

### GET /files/:id

Download file

### DELETE /files/:id

Delete file

### GET /files

List user files

## Audit Endpoints

### GET /audit

Get audit logs (with filters)

### GET /audit/user/:id

Get user activity

### GET /audit/export

Export audit logs

## Health Endpoints

### GET /health

Basic health check

### GET /health/liveness

Liveness probe

### GET /health/readiness

Readiness probe

## Metrics

### GET /metrics

Prometheus metrics endpoint
