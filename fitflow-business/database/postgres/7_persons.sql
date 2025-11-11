-- Person table for FitFlow Business service
-- This table stores common information for all person types (gym owners, trainers, trainees)
-- Each person is linked to a user in the IAM system

-- Create the person table
CREATE TABLE IF NOT EXISTS persons (
    id BIGSERIAL PRIMARY KEY,
    user_id UUID NOT NULL, -- Reference to users table in IAM system
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(30),
    date_of_birth DATE,
    gender VARCHAR(10), -- male, female, other
    profile_image_url VARCHAR(500),
    address VARCHAR(255),
    city VARCHAR(100),
    province VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    emergency_contact_name VARCHAR(150),
    emergency_contact_phone VARCHAR(30),
    emergency_contact_relation VARCHAR(50),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraint to users table (will be handled by application)
    -- CONSTRAINT fk_persons_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_persons_user_id ON persons(user_id);
CREATE INDEX IF NOT EXISTS idx_persons_email ON persons(email);
CREATE INDEX IF NOT EXISTS idx_persons_phone ON persons(phone_number);
CREATE INDEX IF NOT EXISTS idx_persons_is_active ON persons(is_active);
CREATE INDEX IF NOT EXISTS idx_persons_name ON persons(first_name, last_name);
CREATE INDEX IF NOT EXISTS idx_persons_location ON persons(city, province, country);

-- Create trigger for updated_at timestamp
CREATE TRIGGER update_persons_updated_at 
    BEFORE UPDATE ON persons 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE persons IS 'Stores common information for all person types (gym owners, trainers, trainees)';
COMMENT ON COLUMN persons.user_id IS 'Reference to users table in IAM system';
COMMENT ON COLUMN persons.first_name IS 'First name of the person';
COMMENT ON COLUMN persons.last_name IS 'Last name of the person';
COMMENT ON COLUMN persons.email IS 'Email address (unique)';
COMMENT ON COLUMN persons.phone_number IS 'Phone number for contact';
COMMENT ON COLUMN persons.date_of_birth IS 'Date of birth for age calculation';
COMMENT ON COLUMN persons.gender IS 'Gender: male, female, other';
COMMENT ON COLUMN persons.profile_image_url IS 'URL to profile image';
COMMENT ON COLUMN persons.address IS 'Street address';
COMMENT ON COLUMN persons.city IS 'City name';
COMMENT ON COLUMN persons.province IS 'Province/state name';
COMMENT ON COLUMN persons.country IS 'Country name';
COMMENT ON COLUMN persons.postal_code IS 'Postal/ZIP code';
COMMENT ON COLUMN persons.emergency_contact_name IS 'Emergency contact person name';
COMMENT ON COLUMN persons.emergency_contact_phone IS 'Emergency contact phone number';
COMMENT ON COLUMN persons.emergency_contact_relation IS 'Relationship to emergency contact';
COMMENT ON COLUMN persons.is_active IS 'Whether the person is currently active';
