# PostgreSQL Database Scripts

This folder contains PostgreSQL-specific database initialization and schema scripts for the IAM service.

## Files

### `0_init.sql`
PostgreSQL initialization script that runs when the PostgreSQL container starts for the first time.

**What it does:**
- Creates the `iam_db` database
- Creates the `iam_user` user with password `password`
- Grants all privileges on `iam_db` to `iam_user`
- Sets up schema privileges and default privileges
- Configures permissions for future objects

### `1_users_postgres.sql`
User model schema script that creates the users table.

**What it does:**
- Creates the `users` table with all required fields
- Sets up indexes for better performance
- Adds table and column comments for documentation
- Defines constraints and default values
- Creates trigger for automatic timestamp updates

### `2_roles_postgres.sql`
Role model schema script that creates the roles table.

**What it does:**
- Creates the `roles` table with all required fields
- Sets up indexes for better performance
- Adds table and column comments for documentation
- Inserts default roles (user, admin, moderator, guest)
- Creates trigger for automatic timestamp updates

## Usage

### Using Docker Compose
```bash
# Start PostgreSQL with initialization
make docker-up-postgres

# The 0_init.sql script will run automatically
# Then run the schema scripts manually or via GORM migrations
```

### Manual Database Setup
If you prefer to set up the database manually:

```sql
-- Connect to PostgreSQL as postgres user
psql -U postgres

-- Run the initialization script
\i database/postgres/0_init.sql

-- Run the schema scripts
\i database/postgres/1_users_postgres.sql
\i database/postgres/2_roles_postgres.sql
```

### Using GORM Migrations (Recommended)
The application uses GORM AutoMigrate to create tables automatically:

```bash
# Start the application
make run

# GORM will automatically create tables based on Go models
```

## PostgreSQL-Specific Features

### Data Types
- **UUID**: Uses native `UUID` type
- **Timestamps**: Uses `TIMESTAMP` for standard precision
- **Auto-update**: Uses triggers for automatic timestamp updates

### Performance Optimizations
- Indexes on frequently queried fields
- Native UUID type for better performance
- Efficient timestamp handling
- Trigger-based auto-update functionality

### Default Data
- Pre-populated roles: user, admin, moderator, guest
- Uses `ON CONFLICT DO NOTHING` for safe insertion

## Configuration

### Environment Variables
```env
DATABASE_TYPE=postgres
DATABASE_URL=postgres://iam_user:password@localhost:5432/iam_db?sslmode=disable
```

### Connection Parameters
- `sslmode=disable`: Disable SSL for local development
- `sslmode=require`: Require SSL for production

## Advanced Features

### Triggers
- Automatic `updated_at` timestamp updates
- Function-based trigger implementation
- Efficient update handling

### Functions
- `update_updated_at_column()`: Updates timestamp on row changes
- Reusable across multiple tables
- Optimized for performance

## Troubleshooting

### Common Issues
- **Connection refused**: PostgreSQL not running
- **Authentication failed**: Wrong credentials
- **Database does not exist**: Database doesn't exist

### Performance Issues
- **Slow queries**: Check index usage and query plans
- **Memory usage**: Monitor PostgreSQL memory consumption
- **Connection pooling**: Use connection pooling for better performance

## Security Considerations

### Password Storage
- Passwords are hashed using bcrypt
- Password field is excluded from JSON responses

### Role Security
- Role field is hidden from API responses
- Role validation happens server-side
- Default roles provide baseline security

### SSL Configuration
- Use SSL in production environments
- Configure proper SSL certificates
- Monitor SSL connection status

## Migration from MySQL

### Data Migration
1. Export data from MySQL
2. Convert UUID format (CHAR(36) → UUID)
3. Import data to PostgreSQL
4. Update application configuration

### Schema Differences
- UUID type: CHAR(36) → UUID
- Timestamps: DATETIME(3) → TIMESTAMP
- Auto-update: ON UPDATE → Triggers
- Default data: INSERT IGNORE → ON CONFLICT DO NOTHING
