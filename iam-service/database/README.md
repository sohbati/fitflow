# Database Folder

This folder contains database initialization scripts and schema files for the IAM service, organized by database type.

## Folder Structure

```
database/
├── mysql/                   # MySQL-specific scripts
│   ├── 0_init.sql          # MySQL initialization script
│   ├── 1_users.sql         # MySQL users table schema (normalized)
│   ├── 2_roles.sql         # MySQL roles table schema
│   ├── 3_auth_providers.sql # MySQL auth providers table schema
│   ├── 4_user_auth.sql     # MySQL user auth table schema
│   └── README.md           # MySQL documentation
├── postgres/               # PostgreSQL-specific scripts
│   ├── 0_init.sql          # PostgreSQL initialization script
│   ├── 1_users.sql         # PostgreSQL users table schema (normalized)
│   ├── 2_roles.sql         # PostgreSQL roles table schema
│   ├── 3_auth_providers.sql # PostgreSQL auth providers table schema
│   ├── 4_user_auth.sql     # PostgreSQL user auth table schema
│   └── README.md           # PostgreSQL documentation
├── README.md               # This file
└── SCHEMA_OVERVIEW.md      # Schema comparison guide
```

## Quick Start

### MySQL Setup
```bash
# Start MySQL with Docker Compose
make docker-up

# The mysql/0_init.sql script will run automatically
# Tables are created using manual SQL migrations (AutoMigrate disabled)
# Schema files are for reference - GORM handles table creation
```

### PostgreSQL Setup
```bash
# Start PostgreSQL with Docker Compose
make docker-up-postgres

# Update .env file
echo "DATABASE_TYPE=postgres" >> .env
echo "DATABASE_URL=postgres://iam_user:password@localhost:5432/iam_db?sslmode=disable" >> .env

# Tables are created using manual SQL migrations (AutoMigrate disabled)
# Schema files are for reference - GORM handles table creation
```

## Database-Specific Documentation

### MySQL
- **Folder**: `mysql/`
- **Documentation**: `mysql/README.md`
- **Features**: CHAR(36) UUIDs, DATETIME(3) timestamps, MySQL-specific optimizations

### PostgreSQL
- **Folder**: `postgres/`
- **Documentation**: `postgres/README.md`
- **Features**: Native UUID type, TIMESTAMP, trigger-based auto-updates

## Schema Files

### MySQL Scripts
- `mysql/0_init.sql` - Database and user initialization
- `mysql/1_users.sql` - Users table schema
- `mysql/2_roles.sql` - Roles table schema

### PostgreSQL Scripts
- `postgres/0_init.sql` - Database and user initialization
- `postgres/1_users_postgres.sql` - Users table schema
- `postgres/2_roles_postgres.sql` - Roles table schema

## Usage Examples

### Manual MySQL Setup
```sql
-- Connect to MySQL
mysql -u root -p

-- Run initialization
source database/mysql/0_init.sql

-- Run schemas
source database/mysql/1_users.sql
source database/mysql/2_roles.sql
```

### Manual PostgreSQL Setup
```sql
-- Connect to PostgreSQL
psql -U postgres

-- Run initialization
\i database/postgres/0_init.sql

-- Run schemas
\i database/postgres/1_users_postgres.sql
\i database/postgres/2_roles_postgres.sql
```

## Manual SQL Migrations (Recommended)

The application uses manual SQL migrations (AutoMigrate disabled):

```bash
# Start the application
make run

# GORM will automatically create tables based on Go models
# No manual schema execution needed
```

## Configuration

### MySQL Configuration
```env
DATABASE_TYPE=mysql
DATABASE_URL=iam_user:password@tcp(localhost:3306)/iam_db?charset=utf8mb4&parseTime=True&loc=Local
```

### PostgreSQL Configuration
```env
DATABASE_TYPE=postgres
DATABASE_URL=postgres://iam_user:password@localhost:5432/iam_db?sslmode=disable
```

## Key Differences

### MySQL
- Uses `CHAR(36)` for UUID fields
- Uses `DATETIME(3)` for timestamps
- Uses `ON UPDATE CURRENT_TIMESTAMP(3)` for auto-update
- Uses `INSERT IGNORE` for default data

### PostgreSQL
- Uses native `UUID` type
- Uses `TIMESTAMP` for timestamps
- Uses triggers for auto-update functionality
- Uses `ON CONFLICT DO NOTHING` for default data

## Troubleshooting

### Common Issues
- **Connection refused**: Database not running
- **Authentication failed**: Wrong credentials
- **Unknown database**: Database doesn't exist

### Performance Issues
- **Slow queries**: Check index usage
- **Memory usage**: Monitor database memory consumption
- **Connection pooling**: Use connection pooling for better performance

## Security Considerations

### Password Storage
- Passwords are hashed using bcrypt
- Password field is excluded from JSON responses

### Role Security
- Role field is hidden from API responses
- Role validation happens server-side
- Default roles provide baseline security

## Environment Variables

The database connection uses these environment variables:
- `DATABASE_URL`: Connection string
- `DATABASE_TYPE`: Database type (`mysql` or `postgres`)

See `config.example` for example configurations.
