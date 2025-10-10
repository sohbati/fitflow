-- PostgreSQL Role Model Schema
-- This script creates the roles table for the IAM service

-- Create the roles table
CREATE TABLE IF NOT EXISTS roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    role VARCHAR UNIQUE NOT NULL,
    description VARCHAR NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_roles_role ON roles(role);
CREATE INDEX IF NOT EXISTS idx_roles_created_at ON roles(created_at);

-- Add comments for documentation
COMMENT ON TABLE roles IS 'Role definitions for the IAM service';

-- Column comments
COMMENT ON COLUMN roles.id IS 'Unique role identifier (UUID)';
COMMENT ON COLUMN roles.role IS 'Unique role name';
COMMENT ON COLUMN roles.description IS 'Role description';
COMMENT ON COLUMN roles.created_at IS 'Role creation timestamp';
COMMENT ON COLUMN roles.updated_at IS 'Last update timestamp';

-- Create trigger to automatically update updated_at
CREATE TRIGGER update_roles_updated_at 
    BEFORE UPDATE ON roles 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Insert default roles
INSERT INTO roles (role, description) VALUES
('user', 'Standard user with basic permissions'),
('admin', 'Administrator with elevated permissions'),
('moderator', 'Moderator with content management permissions'),
('guest', 'Guest user with limited permissions')
ON CONFLICT (role) DO NOTHING;
