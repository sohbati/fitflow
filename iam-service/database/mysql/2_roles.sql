-- MySQL Role Model Schema
-- This script creates the roles table for the IAM service

-- Create the roles table
drop table if exists roles;
CREATE TABLE IF NOT EXISTS roles (
    id CHAR(36) PRIMARY KEY,
    role VARCHAR(255) UNIQUE NOT NULL,
    description VARCHAR(500) NOT NULL,
    created_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3),
    updated_at DATETIME(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)
);

-- Create indexes for better performance
CREATE INDEX idx_roles_role ON roles(role);
CREATE INDEX idx_roles_created_at ON roles(created_at);

-- Add comments for documentation
ALTER TABLE roles COMMENT = 'Role definitions for the IAM service';

-- Column comments
ALTER TABLE roles MODIFY COLUMN id CHAR(36) COMMENT 'Unique role identifier (UUID)';
ALTER TABLE roles MODIFY COLUMN role VARCHAR(255) COMMENT 'Unique role name';
ALTER TABLE roles MODIFY COLUMN description VARCHAR(500) COMMENT 'Role description';
ALTER TABLE roles MODIFY COLUMN created_at DATETIME(3) COMMENT 'Role creation timestamp';
ALTER TABLE roles MODIFY COLUMN updated_at DATETIME(3) COMMENT 'Last update timestamp';

-- Insert default roles
INSERT IGNORE INTO roles (id, role, description) VALUES
(UUID(), 'user', 'Standard user with basic permissions'),
(UUID(), 'admin', 'Administrator with elevated permissions'),
(UUID(), 'moderator', 'Moderator with content management permissions'),
(UUID(), 'guest', 'Guest user with limited permissions');
