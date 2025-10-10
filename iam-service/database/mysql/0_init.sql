-- MySQL initialization script for IAM service
-- This script runs when the MySQL container starts for the first time

-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS fitflow_iam_db;

-- Use the database
USE fitflow_iam_db;

-- Create the user if it doesn't exist
CREATE USER IF NOT EXISTS 'fitflow_iam_user'@'%' IDENTIFIED BY 'password';

-- Grant all privileges on the database to the user
GRANT ALL PRIVILEGES ON fitflow_iam_db.* TO 'fitflow_iam_user'@'%';

-- Flush privileges to ensure changes take effect
FLUSH PRIVILEGES;

-- Note: Tables will be created automatically by GORM migrations
