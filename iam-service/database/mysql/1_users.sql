-- MySQL User Model Schema (Normalized)
-- This script creates the users table for the IAM service

-- Create the users table
CREATE TABLE IF NOT EXISTS users (
    id CHAR(36) PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    display_name VARCHAR(255) NOT NULL,
    avatar_url VARCHAR(500),
    country VARCHAR(255),
    role VARCHAR(255) DEFAULT 'user',
    is_active BOOLEAN DEFAULT TRUE,
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
);

-- Create indexes for better performance
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_active ON users(is_active);
CREATE INDEX idx_users_role ON users(role);
CREATE INDEX idx_users_created_at ON users(created_at);

-- Add comments for documentation
ALTER TABLE users COMMENT = 'User accounts for the IAM service (normalized)';

-- Column comments
ALTER TABLE users MODIFY COLUMN id CHAR(36) COMMENT 'Unique user identifier (UUID)';
ALTER TABLE users MODIFY COLUMN email VARCHAR(255) COMMENT 'Unique email address';
ALTER TABLE users MODIFY COLUMN display_name VARCHAR(255) COMMENT 'User display name';
ALTER TABLE users MODIFY COLUMN avatar_url VARCHAR(500) COMMENT 'User avatar/profile image URL';
ALTER TABLE users MODIFY COLUMN country VARCHAR(255) COMMENT 'ISO country code';
ALTER TABLE users MODIFY COLUMN role VARCHAR(255) COMMENT 'User role (default: user)';
ALTER TABLE users MODIFY COLUMN is_active BOOLEAN COMMENT 'Whether the user account is active';
ALTER TABLE users MODIFY COLUMN created_at DATETIME(3) COMMENT 'Account creation timestamp';
ALTER TABLE users MODIFY COLUMN updated_at DATETIME(3) COMMENT 'Last update timestamp';
