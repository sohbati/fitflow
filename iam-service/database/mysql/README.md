# MySQL Database Scripts

This folder contains MySQL-specific database initialization and schema scripts for the IAM service.

## Files

### `0_init.sql`
MySQL initialization script that runs when the MySQL container starts for the first time.

**What it does:**
- Creates the `iam_db` database if it doesn't exist
- Creates the `iam_user` user with password `password`
- Grants all privileges on `iam_db` to `iam_user`
- Flushes privileges to ensure changes take effect

### `1_users.sql`
User model schema script that creates the users table.

**What it does:**
- Creates the `users` table with all required fields
- Sets up indexes for better performance
- Adds table and column comments for documentation
- Defines constraints and default values

### `2_roles.sql`
Role model schema script that creates the roles table.

**What it does:**
- Creates the `roles` table with all required fields
- Sets up indexes for better performance
- Adds table and column comments for documentation
- Inserts default roles (user, admin, moderator, guest)

## Usage

### Using Docker Compose
```bash
# Start MySQL with initialization
make docker-up

# The 0_init.sql script will run automatically
# Then run the schema scripts manually or via GORM migrations
```

### Manual Database Setup
If you prefer to set up the database manually:

```sql
-- Connect to MySQL as root
mysql -u root -p

-- Run the initialization script
source database/mysql/0_init.sql

-- Run the schema scripts
source database/mysql/1_users.sql
source database/mysql/2_roles.sql
```

### Using GORM Migrations (Recommended)
The application uses GORM AutoMigrate to create tables automatically:

```bash
# Start the application
make run

# GORM will automatically create tables based on Go models
```

## MySQL-Specific Features

### Data Types
- **UUID**: Uses `CHAR(36)` for UUID fields
- **Timestamps**: Uses `DATETIME(3)` for millisecond precision
- **Auto-update**: Uses `ON UPDATE CURRENT_TIMESTAMP(3)`

### Performance Optimizations
- Indexes on frequently queried fields
- Proper data types for MySQL
- Efficient UUID storage
- Optimized timestamp handling

### Default Data
- Pre-populated roles: user, admin, moderator, guest
- Uses `INSERT IGNORE` for safe insertion

## Configuration

### Environment Variables
```env
DATABASE_TYPE=mysql
DATABASE_URL=iam_user:password@tcp(localhost:3306)/iam_db?charset=utf8mb4&parseTime=True&loc=Local
```

### Connection Parameters
- `charset=utf8mb4`: Character set for MySQL
- `parseTime=True`: Parse time values
- `loc=Local`: Timezone location

## Troubleshooting

### Common Issues
- **Connection refused**: MySQL not running
- **Access denied**: Wrong credentials
- **Unknown database**: Database doesn't exist

### Performance Issues
- **Slow queries**: Check index usage
- **Memory usage**: Monitor MySQL memory consumption
- **Connection pooling**: Use connection pooling for better performance

## Security Considerations

### Password Storage
- Passwords are hashed using bcrypt
- Password field is excluded from JSON responses

### Role Security
- Role field is hidden from API responses
- Role validation happens server-side
- Default roles provide baseline security
