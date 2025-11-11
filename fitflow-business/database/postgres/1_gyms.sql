-- Gym table schema for FitFlow Business service
-- This script creates the main gyms table

-- Create function to update updated_at timestamp (if not exists)
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Create gyms table
CREATE TABLE IF NOT EXISTS gyms (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    phone_number VARCHAR(30),
    email VARCHAR(100),
    website_url VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    facilities JSONB,  -- e.g., {"wifi": true, "pool": false, "sauna": true}
    images JSONB,      -- e.g., [{"url": "...", "type": "logo", "is_primary": true}]
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_gyms_name ON gyms(name);
CREATE INDEX IF NOT EXISTS idx_gyms_verified ON gyms(is_verified);
CREATE INDEX IF NOT EXISTS idx_gyms_created_at ON gyms(created_at);

-- Create trigger for updated_at
CREATE TRIGGER update_gyms_updated_at 
    BEFORE UPDATE ON gyms 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE gyms IS 'Gym information and details';
COMMENT ON COLUMN gyms.id IS 'Unique gym identifier';
COMMENT ON COLUMN gyms.name IS 'Gym name';
COMMENT ON COLUMN gyms.description IS 'Gym description';
COMMENT ON COLUMN gyms.phone_number IS 'Gym contact phone number';
COMMENT ON COLUMN gyms.email IS 'Gym contact email';
COMMENT ON COLUMN gyms.website_url IS 'Gym website URL';
COMMENT ON COLUMN gyms.is_verified IS 'Whether the gym is verified';
COMMENT ON COLUMN gyms.facilities IS 'Gym facilities in JSON format';
COMMENT ON COLUMN gyms.images IS 'Gym images in JSON format';
COMMENT ON COLUMN gyms.created_at IS 'Gym creation timestamp';
COMMENT ON COLUMN gyms.updated_at IS 'Last update timestamp';