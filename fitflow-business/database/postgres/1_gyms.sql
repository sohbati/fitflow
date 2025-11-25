-- Create gyms table inside fitflow_business_schema
CREATE TABLE IF NOT EXISTS fitflow_business_schema.gyms (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description TEXT,
    phone_number VARCHAR(30),
    email VARCHAR(100),
    website_url VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    facilities JSONB,
    images JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX IF NOT EXISTS idx_gyms_name 
    ON fitflow_business_schema.gyms(name);

CREATE INDEX IF NOT EXISTS idx_gyms_verified 
    ON fitflow_business_schema.gyms(is_verified);

CREATE INDEX IF NOT EXISTS idx_gyms_created_at 
    ON fitflow_business_schema.gyms(created_at);

-- Add comments
COMMENT ON TABLE fitflow_business_schema.gyms IS 'Gym information and details';
COMMENT ON COLUMN fitflow_business_schema.gyms.id IS 'Unique gym identifier';
COMMENT ON COLUMN fitflow_business_schema.gyms.name IS 'Gym name';
COMMENT ON COLUMN fitflow_business_schema.gyms.description IS 'Gym description';
COMMENT ON COLUMN fitflow_business_schema.gyms.phone_number IS 'Gym contact phone number';
COMMENT ON COLUMN fitflow_business_schema.gyms.email IS 'Gym contact email';
COMMENT ON COLUMN fitflow_business_schema.gyms.website_url IS 'Gym website URL';
COMMENT ON COLUMN fitflow_business_schema.gyms.is_verified IS 'Whether the gym is verified';
COMMENT ON COLUMN fitflow_business_schema.gyms.facilities IS 'Gym facilities in JSON format';
COMMENT ON COLUMN fitflow_business_schema.gyms.images IS 'Gym images in JSON format';
COMMENT ON COLUMN fitflow_business_schema.gyms.created_at IS 'Gym creation timestamp';
COMMENT ON COLUMN fitflow_business_schema.gyms.updated_at IS 'Last update timestamp';
