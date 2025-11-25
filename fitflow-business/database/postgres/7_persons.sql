-- Person table for FitFlow Business service
-- Stores common information for all person types (gym owners, trainers, trainees)

CREATE TABLE IF NOT EXISTS fitflow_business_schema.persons (
    id BIGSERIAL PRIMARY KEY,

    user_id UUID NOT NULL, -- Reference to IAM users table

    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,

    email VARCHAR(100) UNIQUE,
    phone_number VARCHAR(30),

    date_of_birth DATE,
    gender VARCHAR(10),

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
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_persons_user_id 
    ON fitflow_business_schema.persons(user_id);

CREATE INDEX IF NOT EXISTS idx_persons_email 
    ON fitflow_business_schema.persons(email);

CREATE INDEX IF NOT EXISTS idx_persons_phone 
    ON fitflow_business_schema.persons(phone_number);

CREATE INDEX IF NOT EXISTS idx_persons_is_active 
    ON fitflow_business_schema.persons(is_active);

CREATE INDEX IF NOT EXISTS idx_persons_name 
    ON fitflow_business_schema.persons(first_name, last_name);

CREATE INDEX IF NOT EXISTS idx_persons_location 
    ON fitflow_business_schema.persons(city, province, country);

-- Comments
COMMENT ON TABLE fitflow_business_schema.persons IS 'Stores common information for all person types (gym owners, trainers, trainees)';
COMMENT ON COLUMN fitflow_business_schema.persons.user_id IS 'Reference to IAM users table';
COMMENT ON COLUMN fitflow_business_schema.persons.first_name IS 'First name';
COMMENT ON COLUMN fitflow_business_schema.persons.last_name IS 'Last name';
COMMENT ON COLUMN fitflow_business_schema.persons.email IS 'Unique email address';
COMMENT ON COLUMN fitflow_business_schema.persons.phone_number IS 'Phone number';
COMMENT ON COLUMN fitflow_business_schema.persons.profile_image_url IS 'URL to profile image';
COMMENT ON COLUMN fitflow_business_schema.persons.is_active IS 'Whether the person account is active';
