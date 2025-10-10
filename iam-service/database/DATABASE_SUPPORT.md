# Database Support Documentation

This document describes the multi-database support implemented in the IAM service, with MySQL as the default database and PostgreSQL as an alternative option.

## Supported Databases

### 1. MySQL (Default)
- **Version**: MySQL 8.0+
- **Driver**: `gorm.io/driver/mysql`
- **Default Port**: 3306
- **Default Database**: `iam_db`
- **Default User**: `iam_user`

### 2. PostgreSQL
- **Version**: PostgreSQL 15+
- **Driver**: `gorm.io/driver/postgres`
- **Default Port**: 5432
- **Default Database**: `iam_db`
- **Default User**: `iam_user`

## Configuration

### Environment Variables

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `DATABASE_TYPE` | Database type | `mysql` | `mysql`, `postgres` |
| `DATABASE_URL` | Database connection string | MySQL default | See examples below |

### MySQL Configuration

#### Default Configuration
```env
DATABASE_TYPE=mysql
DATABASE_URL=iam_user:password@tcp(localhost:3306)/iam_db?charset=utf8mb4&parseTime=True&loc=Local
```

#### Connection String Format
```
username:password@tcp(host:port)/database?charset=utf8mb4&parseTime=True&loc=Local
```

#### Parameters
- `charset=utf8mb4`: Character set for MySQL
- `parseTime=True`: Parse time values
- `loc=Local`: Timezone location

### PostgreSQL Configuration

#### Configuration
```env
DATABASE_TYPE=postgres
DATABASE_URL=postgres://iam_user:password@localhost:5432/iam_db?sslmode=disable
```

#### Connection String Format
```
postgres://username:password@host:port/database?sslmode=disable
```

#### Parameters
- `sslmode=disable`: Disable SSL for local development

## Implementation Details

### Database Connection Logic

```go
func Connect(databaseURL, databaseType string) (*gorm.DB, error) {
    var db *gorm.DB
    var err error

    switch databaseType {
    case "mysql":
        db, err = gorm.Open(mysql.Open(databaseURL), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        })
    case "postgres":
        db, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
            Logger: logger.Default.LogMode(logger.Info),
        })
    default:
        log.Fatalf("Unsupported database type: %s. Supported types: mysql, postgres", databaseType)
    }

    if err != nil {
        return nil, err
    }

    log.Printf("Database connected successfully using %s", databaseType)
    return db, nil
}
```

### Configuration Structure

```go
type Config struct {
    JWTSecret        string
    DatabaseURL      string
    DatabaseType     string  // NEW: Database type field
    TokenExpMinutes  int
    Port             string
    TokenExpDuration time.Duration
}
```

## Docker Compose Setup

### MySQL Service (Default)
```yaml
mysql:
  image: mysql:8.0
  container_name: iam-mysql
  restart: unless-stopped
  environment:
    MYSQL_ROOT_PASSWORD: rootpassword
    MYSQL_DATABASE: iam_db
    MYSQL_USER: iam_user
    MYSQL_PASSWORD: password
  ports:
    - "3306:3306"
  volumes:
    - mysql_data:/var/lib/mysql
    - ./database/init.sql:/docker-entrypoint-initdb.d/init.sql
  command: --default-authentication-plugin=mysql_native_password
```

### PostgreSQL Service (Optional)
```yaml
postgres:
  image: postgres:15
  container_name: iam-postgres
  restart: unless-stopped
  environment:
    POSTGRES_DB: iam_db
    POSTGRES_USER: iam_user
    POSTGRES_PASSWORD: password
  ports:
    - "5432:5432"
  volumes:
    - postgres_data:/var/lib/postgresql/data
  profiles:
    - postgres
```

## Makefile Commands

### MySQL Commands
```bash
# Start MySQL database
make docker-up

# Development environment with MySQL
make dev
```

### PostgreSQL Commands
```bash
# Start PostgreSQL database
make docker-up-postgres

# Development environment with PostgreSQL
make dev-postgres
```

### General Commands
```bash
# Stop all services
make docker-down

# Complete setup
make setup
```

## Database Schema

### User Table
```sql
-- MySQL
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    mobile VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    country VARCHAR(255) NOT NULL,
    role VARCHAR(255) DEFAULT 'user',
    created_at DATETIME(3),
    updated_at DATETIME(3)
);

-- PostgreSQL
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR UNIQUE NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    mobile VARCHAR UNIQUE NOT NULL,
    display_name VARCHAR NOT NULL,
    password VARCHAR NOT NULL,
    country VARCHAR NOT NULL,
    role VARCHAR DEFAULT 'user',
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

### Role Table
```sql
-- MySQL
CREATE TABLE roles (
    id CHAR(36) PRIMARY KEY,
    role VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(255) NOT NULL,
    created_at DATETIME(3),
    updated_at DATETIME(3)
);

-- PostgreSQL
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);
```

## Migration Differences

### MySQL
- Uses `CHAR(36)` for UUID fields
- Uses `DATETIME(3)` for timestamps
- Auto-increment for primary keys (if not using UUID)

### PostgreSQL
- Uses `UUID` type for UUID fields
- Uses `TIMESTAMP` for timestamps
- Uses `gen_random_uuid()` for UUID generation

## GORM Compatibility

### Model Definitions
```go
type User struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Username    string    `gorm:"uniqueIndex;not null" json:"username"`
    Email       string    `gorm:"uniqueIndex;not null" json:"email"`
    Mobile      string    `gorm:"uniqueIndex;not null" json:"mobile"`
    DisplayName string    `gorm:"not null" json:"display_name"`
    Password    string    `gorm:"not null" json:"-"`
    Country     string    `gorm:"not null" json:"country"`
    Role        string    `gorm:"default:user" json:"-"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

### GORM Tags
- `type:uuid`: Specifies UUID type (PostgreSQL)
- `primary_key`: Primary key constraint
- `default:gen_random_uuid()`: Default UUID generation (PostgreSQL)
- `uniqueIndex`: Unique index constraint
- `not null`: Not null constraint
- `default:user`: Default value

## Performance Considerations

### MySQL
- **Advantages**: Fast for simple queries, good for read-heavy workloads
- **Considerations**: UUID performance, character set optimization

### PostgreSQL
- **Advantages**: Advanced features, JSON support, better for complex queries
- **Considerations**: Memory usage, connection pooling

## Security Considerations

### Connection Security
- **MySQL**: Use SSL in production (`tls=true`)
- **PostgreSQL**: Use SSL in production (`sslmode=require`)

### Authentication
- **MySQL**: Use strong passwords, limit user privileges
- **PostgreSQL**: Use strong passwords, limit user privileges

## Troubleshooting

### Common Issues

#### MySQL Connection Issues
```bash
# Check if MySQL is running
docker ps | grep mysql

# Check MySQL logs
docker logs iam-mysql

# Test connection
mysql -h localhost -P 3306 -u iam_user -p iam_db
```

#### PostgreSQL Connection Issues
```bash
# Check if PostgreSQL is running
docker ps | grep postgres

# Check PostgreSQL logs
docker logs iam-postgres

# Test connection
psql -h localhost -p 5432 -U iam_user -d iam_db
```

### Error Messages

#### MySQL Errors
- `dial tcp 127.0.0.1:3306: connect: connection refused`: MySQL not running
- `Access denied for user`: Authentication failed
- `Unknown database`: Database doesn't exist

#### PostgreSQL Errors
- `dial tcp 127.0.0.1:5432: connect: connection refused`: PostgreSQL not running
- `authentication failed`: Authentication failed
- `database does not exist`: Database doesn't exist

## Best Practices

### 1. Environment-Specific Configuration
- Use different configurations for development, staging, and production
- Use environment variables for sensitive information
- Use connection pooling in production

### 2. Database Selection
- **Choose MySQL for**: Simple applications, high read performance, cost-sensitive deployments
- **Choose PostgreSQL for**: Complex applications, advanced features, data integrity requirements

### 3. Migration Strategy
- Use GORM AutoMigrate for development
- Use proper migration tools for production
- Test migrations on staging environment first

### 4. Monitoring
- Monitor database performance
- Set up alerts for connection issues
- Use database-specific monitoring tools

## Future Enhancements

### Potential Improvements
1. **Connection Pooling**: Implement connection pooling for better performance
2. **Read Replicas**: Support for read replicas
3. **Database Sharding**: Support for horizontal scaling
4. **Backup Integration**: Automated backup solutions
5. **Performance Monitoring**: Database performance monitoring
6. **Health Checks**: Database health check endpoints

### Additional Database Support
1. **SQLite**: For development and testing
2. **SQL Server**: For enterprise environments
3. **Oracle**: For enterprise environments
4. **CockroachDB**: For distributed databases

This multi-database support provides flexibility while maintaining consistency across different database systems.
