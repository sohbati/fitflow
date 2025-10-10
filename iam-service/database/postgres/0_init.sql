-- PostgreSQL initialization script for IAM service
-- This script runs when the PostgreSQL container starts for the first time

-- Create the database if it doesn't exist
-- Note: PostgreSQL doesn't support CREATE DATABASE IF NOT EXISTS
-- This should be run as the postgres superuser

-- Create the database
CREATE DATABASE fitflow_iam_db;

-- Create the user if it doesn't exist
-- Note: PostgreSQL doesn't support CREATE USER IF NOT EXISTS
-- This should be run as the postgres superuser
CREATE USER fitflow_iam_user WITH PASSWORD 'password';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE fitflow_iam_db TO fitflow_iam_user;

-- Connect to the newly created database
\c fitflow_iam_db;

-- Grant schema privileges
GRANT ALL ON SCHEMA public TO fitflow_iam_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO fitflow_iam_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO fitflow_iam_user;

-- Set default privileges for future objects
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO fitflow_iam_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO fitflow_iam_user;

-- Note: Tables will be created automatically by GORM migrations
