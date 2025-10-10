# Database Schema Overview

This document provides an overview of the database schemas for both MySQL and PostgreSQL implementations.

## File Structure

```
database/
├── mysql/                   # MySQL-specific scripts
│   ├── 0_init.sql          # MySQL initialization script
│   ├── 1_users.sql         # MySQL users table schema
│   ├── 2_roles.sql         # MySQL roles table schema
│   └── README.md           # MySQL documentation
├── postgres/               # PostgreSQL-specific scripts
│   ├── 0_init.sql              # PostgreSQL initialization script
│   ├── 1_users_postgres.sql    # PostgreSQL users table schema
│   ├── 2_roles_postgres.sql    # PostgreSQL roles table schema
│   └── README.md               # PostgreSQL documentation
├── README.md               # Main database documentation
└── SCHEMA_OVERVIEW.md      # This file
```

## Schema Comparison

### Users Table

| Field | MySQL Type | PostgreSQL Type | Description |
|-------|------------|-----------------|-------------|
| `id` | `CHAR(36)` | `UUID` | Primary key |
| `username` | `VARCHAR(255)` | `VARCHAR` | Unique username |
| `email` | `VARCHAR(255)` | `VARCHAR` | Unique email |
| `mobile` | `VARCHAR(255)` | `VARCHAR` | Unique mobile |
| `display_name` | `VARCHAR(255)` | `VARCHAR` | Display name |
| `password` | `VARCHAR(255)` | `VARCHAR` | Hashed password |
| `country` | `VARCHAR(255)` | `VARCHAR` | Country code |
| `role` | `VARCHAR(255)` | `VARCHAR` | User role |
| `created_at` | `DATETIME(3)` | `TIMESTAMP` | Creation time |
| `updated_at` | `DATETIME(3)` | `TIMESTAMP` | Update time |

### Roles Table

| Field | MySQL Type | PostgreSQL Type | Description |
|-------|------------|-----------------|-------------|
| `id` | `CHAR(36)` | `UUID` | Primary key |
| `role` | `VARCHAR(255)` | `VARCHAR` | Unique role name |
| `description` | `VARCHAR(500)` | `VARCHAR` | Role description |
| `created_at` | `DATETIME(3)` | `TIMESTAMP` | Creation time |
| `updated_at` | `DATETIME(3)` | `TIMESTAMP` | Update time |

## Key Differences

### MySQL Implementation
- Uses `CHAR(36)` for UUID fields
- Uses `DATETIME(3)` for timestamps
- Uses `ON UPDATE CURRENT_TIMESTAMP(3)` for auto-update
- Uses `INSERT IGNORE` for default data

### PostgreSQL Implementation
- Uses native `UUID` type
- Uses `TIMESTAMP` for timestamps
- Uses triggers for auto-update functionality
- Uses `ON CONFLICT DO NOTHING` for default data

## Indexes

### Users Table Indexes
- `idx_users_username` - Username lookup
- `idx_users_email` - Email lookup
- `idx_users_mobile` - Mobile lookup
- `idx_users_role` - Role filtering
- `idx_users_created_at` - Date filtering

### Roles Table Indexes
- `idx_roles_role` - Role lookup
- `idx_roles_created_at` - Date filtering

## Default Data

### Default Roles
- `user` - Standard user with basic permissions
- `admin` - Administrator with elevated permissions
- `moderator` - Moderator with content management permissions
- `guest` - Guest user with limited permissions

## Usage Examples

### MySQL Setup
```bash
# Start MySQL container
make docker-up

# Connect to MySQL
mysql -h localhost -P 3306 -u iam_user -p iam_db

# Run schema scripts
source database/mysql/1_users.sql
source database/mysql/2_roles.sql
```

### PostgreSQL Setup
```bash
# Start PostgreSQL container
make docker-up-postgres

# Connect to PostgreSQL
psql -h localhost -p 5432 -U iam_user -d iam_db

# Run initialization script
\i database/postgres/0_init.sql

# Run schema scripts
\i database/postgres/1_users_postgres.sql
\i database/postgres/2_roles_postgres.sql
```

### GORM AutoMigrate (Recommended)
```bash
# Start the application
make run

# GORM will automatically create tables based on Go models
```

## Performance Considerations

### MySQL
- Use `CHAR(36)` for UUIDs (fixed length, better performance)
- Use `DATETIME(3)` for millisecond precision
- Indexes are created for optimal query performance

### PostgreSQL
- Use native `UUID` type for better performance
- Use `TIMESTAMP` for standard precision
- Triggers handle automatic timestamp updates

## Security Considerations

### Password Storage
- Passwords are hashed using bcrypt
- Password field is excluded from JSON responses
- Minimum password requirements should be enforced

### Role Security
- Role field is hidden from API responses
- Role validation happens server-side
- Default roles provide baseline security

## Migration Strategy

### From MySQL to PostgreSQL
1. Export data from MySQL
2. Convert UUID format (CHAR(36) → UUID)
3. Import data to PostgreSQL
4. Update application configuration

### From PostgreSQL to MySQL
1. Export data from PostgreSQL
2. Convert UUID format (UUID → CHAR(36))
3. Import data to MySQL
4. Update application configuration

## Maintenance

### Regular Tasks
- Monitor database performance
- Check index usage
- Update statistics
- Backup data regularly

### Schema Updates
- Use GORM migrations for application updates
- Use manual SQL scripts for complex changes
- Test changes in development first
- Document all schema changes

## Troubleshooting

### Common Issues
- **UUID Format**: Ensure consistent UUID format across databases
- **Timezone**: Set proper timezone for timestamp fields
- **Indexes**: Monitor index usage and performance
- **Constraints**: Verify all constraints are properly defined

### Performance Issues
- **Slow Queries**: Check index usage and query plans
- **Lock Contention**: Monitor for table locks and deadlocks
- **Memory Usage**: Monitor database memory consumption
- **Connection Pooling**: Use connection pooling for better performance

This schema overview provides a comprehensive guide for understanding and maintaining the database schemas across both MySQL and PostgreSQL implementations.
