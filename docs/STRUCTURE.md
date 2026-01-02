# Project Structure Documentation

## 🎯 Design Principles

1. **Single Responsibility**: Each file has one clear purpose
2. **Max 80 Lines**: Files kept small and focused
3. **Clear Separation**: Models, DTOs, Handlers, Services separated
4. **Easy Navigation**: Intuitive file naming and organization

## 📁 Detailed Structure

### 1. Domain Layer (`internal/domain/`)

#### Models (`internal/domain/models/`)

- **`user_model.go`** - User struct definition (30 lines)
- **`user_role.go`** - User role enum and methods (25 lines)
- **`user_methods.go`** - User domain methods (60 lines)
- **`session.go`** - Session model (20 lines)
- **`audit.go`** - Audit log model (30 lines)
- **`file.go`** - File model (25 lines)
- **`token.go`** - Token models (25 lines)

**Benefits:**

- Easy to find specific model
- Clear separation of concerns
- Methods grouped logically

#### Constants (`internal/domain/constants/`)

- **`user_constants.go`** - User-related constants (30 lines)
- **`auth_constants.go`** - Auth constants (35 lines)
- **`file_constants.go`** - File constants (40 lines)
- **`ratelimit_constants.go`** - Rate limit configs (30 lines)
- **`cache_constants.go`** - Cache key constants (25 lines)

**Benefits:**

- Centralized configuration
- Type-safe constants
- Easy to maintain

### 2. API Layer (`internal/api/`)

#### DTOs (`internal/api/dto/`)

- **`auth_dto.go`** - Login/register DTOs (30 lines)
- **`auth_request.go`** - Auth requests (40 lines)
- **`mfa_dto.go`** - MFA DTOs (30 lines)
- **`user_dto.go`** - User DTOs (50 lines)
- **`file_dto.go`** - File DTOs (35 lines)
- **`audit_dto.go`** - Audit DTOs (45 lines)
- **`common_dto.go`** - Common responses (40 lines)

**Benefits:**

- Clear API contracts
- Validation in one place
- Easy to document

#### Handlers (`internal/api/handlers/`)

**Auth Handlers:**

- **`auth_handler.go`** - Base handler (15 lines)
- **`auth_login.go`** - Login logic (60 lines)
- **`auth_register.go`** - Register logic (60 lines)
- **`auth_logout.go`** - Logout logic (40 lines)
- **`auth_refresh.go`** - Token refresh (50 lines)
- **`auth_password.go`** - Password reset (70 lines)
- **`auth_mfa.go`** - MFA logic (70 lines)

**User Handlers:**

- **`user_handler.go`** - Base handler (15 lines)
- **`user_get.go`** - Get user (50 lines)
- **`user_list.go`** - List users (60 lines)
- **`user_update.go`** - Update user (70 lines)
- **`user_delete.go`** - Delete user (50 lines)
- **`user_profile.go`** - Profile management (70 lines)
- **`user_avatar.go`** - Avatar upload (75 lines)

**Benefits:**

- Each endpoint in separate file
- Easy to test individually
- Clear responsibilities

#### Mappers (`internal/api/mapper/`)

- **`user_mapper.go`** - User model ↔ DTO (50 lines)
- **`file_mapper.go`** - File model ↔ DTO (40 lines)
- **`audit_mapper.go`** - Audit model ↔ DTO (35 lines)

**Benefits:**

- Centralized conversion logic
- DRY principle
- Easy to maintain

### 3. Error Handling (`pkg/errors/`)

- **`error_types.go`** - Error types (40 lines)
- **`auth_errors.go`** - Auth errors (45 lines)
- **`user_errors.go`** - User errors (45 lines)
- **`validation_errors.go`** - Validation errors (50 lines)
- **`resource_errors.go`** - Resource errors (35 lines)
- **`system_errors.go`** - System errors (45 lines)
- **`http.go`** - HTTP helpers (30 lines)

**Benefits:**

- Typed errors
- Consistent error handling
- Easy to add new errors

## 📊 File Size Distribution

| File Type | Avg Lines | Max Lines | Status  |
| --------- | --------- | --------- | ------- |
| Models    | 25        | 60        | ✅ Good |
| DTOs      | 35        | 50        | ✅ Good |
| Handlers  | 55        | 75        | ✅ Good |
| Constants | 30        | 40        | ✅ Good |
| Errors    | 40        | 50        | ✅ Good |
| Mappers   | 45        | 50        | ✅ Good |

## 🚀 Future Development Guidelines

### Adding New Features

1. **New Entity (e.g., Product):**

   ```
   internal/domain/models/
   ├── product_model.go      (struct definition)
   ├── product_status.go     (enums/constants)
   └── product_methods.go    (domain methods)

   internal/domain/constants/
   └── product_constants.go  (constants)

   internal/api/dto/
   ├── product_dto.go        (main DTOs)
   └── product_request.go    (request DTOs)

   internal/api/handlers/
   ├── product_handler.go    (base)
   ├── product_list.go       (list endpoint)
   ├── product_get.go        (get endpoint)
   ├── product_create.go     (create endpoint)
   ├── product_update.go     (update endpoint)
   └── product_delete.go     (delete endpoint)

   internal/api/mapper/
   └── product_mapper.go     (conversions)

   pkg/errors/
   └── product_errors.go     (errors)
   ```

2. **New Handler Method:**

   - Create new file: `{entity}_{action}.go`
   - Max 60-80 lines including comments
   - Single responsibility

3. **New DTO:**
   - Add to existing DTO file if related
   - Create new file if different concern
   - Keep under 50 lines

### Code Review Checklist

- [ ] File < 80 lines
- [ ] Single responsibility
- [ ] Clear naming
- [ ] Proper package
- [ ] No duplication
- [ ] Comments for TODO

## 📝 Naming Conventions

### Files

- Models: `{entity}_model.go`, `{entity}_methods.go`
- DTOs: `{entity}_dto.go`, `{entity}_request.go`
- Handlers: `{entity}_{action}.go`
- Constants: `{domain}_constants.go`
- Errors: `{domain}_errors.go`

### Functions

- Handlers: `{Action}{Entity}` (e.g., `CreateUser`)
- Services: `{Action}{Entity}` (e.g., `CreateUser`)
- Mappers: `{Entity}ToDTO`, `DTOTo{Entity}`

## ✅ Benefits of This Structure

1. **Easy Navigation**: Find any code in seconds
2. **Isolated Changes**: Change one feature without affecting others
3. **Parallel Development**: Multiple developers work simultaneously
4. **Testing**: Test individual files/functions easily
5. **Code Review**: Review small, focused changes
6. **Onboarding**: New developers understand quickly
7. **Maintenance**: Fix bugs in specific files
8. **Scalability**: Add features without bloat

## 🎓 Next Steps

1. Implement handlers following template
2. Add service layer with similar separation
3. Add repository layer with query separation
4. Add comprehensive tests per file
5. Document each package
