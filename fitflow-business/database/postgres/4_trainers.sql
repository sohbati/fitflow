-- Create trainers table inside schema
CREATE TABLE IF NOT EXISTS fitflow_business_schema.trainers (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL ,
    is_registered BOOLEAN DEFAULT FALSE, -- registered with a gym or independent
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_trainers_person_id 
    ON fitflow_business_schema.trainers(person_id);

CREATE INDEX IF NOT EXISTS idx_trainers_registered 
    ON fitflow_business_schema.trainers(is_registered);

-- Add comments for documentation
COMMENT ON TABLE fitflow_business_schema.trainers 
    IS 'Trainer information linked to persons table';

COMMENT ON COLUMN fitflow_business_schema.trainers.id 
    IS 'Unique trainer identifier';

COMMENT ON COLUMN fitflow_business_schema.trainers.person_id 
    IS 'Reference to persons table';

COMMENT ON COLUMN fitflow_business_schema.trainers.is_registered 
    IS 'Whether trainer is registered with a gym';

COMMENT ON COLUMN fitflow_business_schema.trainers.created_at 
    IS 'Trainer creation timestamp';

COMMENT ON COLUMN fitflow_business_schema.trainers.updated_at 
    IS 'Last update timestamp';
