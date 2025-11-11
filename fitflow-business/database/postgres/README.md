# FitFlow Business Database Schema Files

This directory contains the PostgreSQL schema files for the FitFlow Business service, organized by table for better maintainability.

## File Structure

### `0_init.sql`
- Database initialization script
- Creates database and user
- Sets up permissions and schema privileges
- **Must be run as postgres superuser**

### `1_gyms.sql`
- Main gyms table
- Contains gym information, facilities (JSONB), and images (JSONB)
- Includes indexes and triggers for performance and automatic timestamps

### `2_gym_locations.sql`
- Gym locations table
- Supports multiple locations per gym
- Includes geographic coordinates (latitude/longitude)
- Foreign key reference to gyms table

### `3_gym_owners.sql`
- Gym owners table
- Supports multiple owners per gym
- Includes owner contact information and bio
- Foreign key reference to gyms table

### `4_trainers.sql`
- Trainers table
- Independent trainers or gym-registered trainers
- Includes trainer contact information
- Can be linked to multiple gyms

### `5_gym_trainers.sql`
- Junction table for gym-trainer relationships
- Many-to-many relationship between gyms and trainers
- Includes unique constraint to prevent duplicates
- Foreign key references to both gyms and trainers tables

## Execution Order

The files are numbered to ensure proper execution order:

1. **0_init.sql** - Database setup (run as superuser)
2. **1_gyms.sql** - Create main gyms table
3. **2_gym_locations.sql** - Create locations table (depends on gyms)
4. **3_gym_owners.sql** - Create owners table (depends on gyms)
5. **4_trainers.sql** - Create trainers table (independent)
6. **5_gym_trainers.sql** - Create junction table (depends on both gyms and trainers)

## Docker Setup

When using Docker Compose, all files are automatically executed in order during container initialization:

```bash
cd fitflow-business
docker-compose up -d postgres
```

## Manual Setup

For manual setup, run the files in order:

```bash
# 1. Run initialization as superuser
psql -U postgres -d postgres -f database/postgres/0_init.sql

# 2. Connect to the business database and run schema files
psql -U fitflow_business_user -d fitflow_business_db -f database/postgres/1_gyms.sql
psql -U fitflow_business_user -d fitflow_business_db -f database/postgres/2_gym_locations.sql
psql -U fitflow_business_user -d fitflow_business_db -f database/postgres/3_gym_owners.sql
psql -U fitflow_business_user -d fitflow_business_db -f database/postgres/4_trainers.sql
psql -U fitflow_business_user -d fitflow_business_db -f database/postgres/5_gym_trainers.sql
```

## Features

- **Modular Design**: Each table in its own file for easy maintenance
- **Proper Dependencies**: Foreign key relationships maintained
- **Performance Optimized**: Indexes on frequently queried columns
- **Automatic Timestamps**: Triggers for updated_at fields
- **Documentation**: Comprehensive comments on tables and columns
- **JSONB Support**: Flexible data storage for facilities and images
- **Geographic Data**: Support for location coordinates
- **Many-to-Many Relationships**: Proper junction table implementation

## Schema Relationships

```
gyms (1) ──→ (N) gym_locations
gyms (1) ──→ (N) gym_owners
gyms (N) ←──→ (N) trainers (via gym_trainers)
```

This modular approach makes it easier to:
- Add new tables without affecting existing ones
- Modify individual tables independently
- Understand the database structure
- Maintain and version control schema changes
