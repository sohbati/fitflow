# Database Setup Guide

## Overview

This project uses manual SQL migrations for both development and production. GORM AutoMigrate is disabled to ensure consistent database schema across all environments.

## Database Schema

### 1. Users Table (Normalized)
```sql
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(500),
    country VARCHAR(255),
    role VARCHAR(255) DEFAULT 'user',
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 2. Roles Table
```sql
CREATE TABLE roles (
    id CHAR(36) PRIMARY KEY,
    role VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(500),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 3. Auth Providers Table
```sql
CREATE TABLE auth_providers (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL,
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

### 4. User Auth Table
```sql
CREATE TABLE user_auth (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    auth_provider_id CHAR(36) NOT NULL,
    provider_user_id VARCHAR(255),
    provider_email VARCHAR(255),
    username VARCHAR(255),
    provider_data JSON,
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP NULL,
    otp_code VARCHAR(10) NULL,
    otp_expires_at TIMESTAMP NULL,
    otp_attempts INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (auth_provider_id) REFERENCES auth_providers(id) ON DELETE CASCADE,
    UNIQUE KEY unique_user_provider (user_id, auth_provider_id),
    UNIQUE KEY unique_provider_user_id (auth_provider_id, provider_user_id)
);
```

## Setup Instructions

### 1. Start Database with Docker Compose

#### MySQL
```bash
# Start MySQL
docker-compose up -d mysql

# Or use the make command
make docker-up
```

#### PostgreSQL
```bash
# Start PostgreSQL
docker-compose up -d postgres

# Or use the make command
make docker-up-postgres
```

### 2. Manual Database Setup (Optional)

If you prefer to set up the database manually, you can run the individual schema files:

#### MySQL
```bash
# Connect to MySQL
mysql -u fitflow_iam_user -p fitflow_iam_db

# Run the schema files in order
source database/mysql/1_users.sql
source database/mysql/2_roles.sql
source database/mysql/3_auth_providers.sql
source database/mysql/4_user_auth.sql
```

#### PostgreSQL
```bash
# Connect to PostgreSQL
psql -U fitflow_iam_user -d fitflow_iam_db

# Run the schema files in order
\i database/postgres/1_users.sql
\i database/postgres/2_roles.sql
\i database/postgres/3_auth_providers.sql
\i database/postgres/4_user_auth.sql
```

### 3. Configure Environment Variables

Create a `.env` file:
```env
# Database Configuration
DATABASE_TYPE=mysql
DATABASE_URL=fitflow_iam_user:password@tcp(localhost:3306)/fitflow_iam_db?charset=utf8mb4&parseTime=True&loc=Local

# JWT Configuration
JWT_SECRET=your-secret-key
TOKEN_EXP_MINUTES=15

# Server Configuration
PORT=8091

# Google OAuth (Optional)
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
GOOGLE_REDIRECT_URL=http://localhost:3000/auth/google/callback
```

### 4. Run the Application

```bash
# Build the application
go build -o iam-service .

# Run the application
./iam-service
```

The application will automatically:
- Connect to the database
- Use manual SQL migrations (AutoMigrate disabled)
- Initialize default auth providers
- Start the HTTP server

## Manual SQL Migrations

The application uses manual SQL migrations for database schema management:

```go
func Migrate(db *gorm.DB) error {
    // AutoMigrate disabled for both development and production
    // Use manual SQL migrations instead
    log.Println("AutoMigrate disabled - using manual SQL migrations")
    return nil
}
```

### Benefits of Manual Migrations
- **Consistent Schema**: Same schema across all environments
- **Version Control**: SQL scripts are tracked in git
- **Production Safety**: No unexpected schema changes
- **Explicit Control**: Full control over database structure

### Migration Process
1. **Development**: Run SQL scripts manually when needed
2. **Production**: Apply migrations during deployment
3. **Version Control**: All schema changes tracked in git
4. **Rollback**: Manual rollback scripts if needed

## Default Data

### Auth Providers
The application automatically creates default auth providers:
- `local` - Username/password authentication
- `google` - Google OAuth 2.0
- `facebook` - Facebook OAuth 2.0 (ready for implementation)
- `apple` - Apple Sign In (ready for implementation)
- `mobile_otp` - Mobile phone OTP (ready for implementation)
- `email_otp` - Email OTP (ready for implementation)

### Roles
Default roles can be created manually:
- `user` - Regular user (default)
- `admin` - Administrator
- `moderator` - Moderator

## Schema Files

The schema files in the `database/` folder are for reference only:
- `mysql/1_users.sql` - MySQL users table schema
- `mysql/2_roles.sql` - MySQL roles table schema
- `mysql/3_auth_providers.sql` - MySQL auth providers table schema
- `mysql/4_user_auth.sql` - MySQL user auth table schema
- `postgres/1_users.sql` - PostgreSQL users table schema
- `postgres/2_roles.sql` - PostgreSQL roles table schema
- `postgres/3_auth_providers.sql` - PostgreSQL auth providers table schema
- `postgres/4_user_auth.sql` - PostgreSQL user auth table schema

## Database URLs

### MySQL
```
DATABASE_URL=fitflow_iam_user:password@tcp(localhost:3306)/fitflow_iam_db?charset=utf8mb4&parseTime=True&loc=Local
```

### PostgreSQL
```
DATABASE_URL=postgres://fitflow_iam_user:password@localhost:5432/fitflow_iam_db?sslmode=disable
```

## Troubleshooting

### Common Issues

1. **Connection Refused**
   - Ensure database is running: `docker-compose ps`
   - Check database URL in `.env` file
   - Verify database credentials

2. **Table Creation Failed**
   - Check database permissions
   - Ensure database exists
   - Review GORM logs

3. **Migration Errors**
   - Stop application
   - Drop database and recreate
   - Restart application

### Reset Database

```bash
# Stop application
# Drop database
docker-compose down -v

# Restart database
docker-compose up -d

# Restart application
./iam-service
```

## Production Considerations

For production deployment:

1. **Use Proper Migrations**: Implement proper database migration system
2. **Backup Strategy**: Implement database backup and recovery
3. **Connection Pooling**: Configure database connection pooling
4. **Monitoring**: Add database monitoring and alerting
5. **Security**: Use secure database credentials and connections

## Next Steps

1. **Add New Auth Providers**: Follow the pattern in `GoogleAuthHandler`
2. **Implement OTP**: Add SMS/Email service integration
3. **Add Validation**: Implement input validation and sanitization
4. **Add Logging**: Implement structured logging
5. **Add Tests**: Write unit and integration tests
