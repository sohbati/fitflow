-- PostgreSQL initialization script for FitFlow Business service
-- IMPORTANT: This script must be run as the postgres superuser

-- Create the database if it doesn't exist
CREATE DATABASE fitflow_business_db;

-- Create the user if it doesn't exist
CREATE USER fitflow_business_user WITH PASSWORD 'password';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON DATABASE fitflow_business_db TO fitflow_business_user;

-- Connect to the newly created database
\c fitflow_business_db;

-- Grant schema privileges (this is crucial!)
GRANT ALL ON SCHEMA public TO fitflow_business_user;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO fitflow_business_user;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO fitflow_business_user;

-- Set default privileges for future objects
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON TABLES TO fitflow_business_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT ALL ON SEQUENCES TO fitflow_business_user;
 
 CREATE SCHEMA fitflow_business_schema AUTHORIZATION fitflow_business_user;
