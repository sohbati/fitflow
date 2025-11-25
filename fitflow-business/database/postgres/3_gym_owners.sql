-- Create gym_owners table inside schema
CREATE TABLE IF NOT EXISTS fitflow_business_schema.gym_owners (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL ,
    gym_id BIGINT NOT NULL ,
    brief_bio TEXT,    -- short description about the owner
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_gym_owners_person_id 
    ON fitflow_business_schema.gym_owners(person_id);

CREATE INDEX IF NOT EXISTS idx_gym_owners_gym_id 
    ON fitflow_business_schema.gym_owners(gym_id);

CREATE INDEX IF NOT EXISTS idx_gym_owners_person_gym 
    ON fitflow_business_schema.gym_owners(person_id, gym_id);

-- Add comments for documentation
COMMENT ON TABLE fitflow_business_schema.gym_owners 
    IS 'Gym owner information linked to persons table';

COMMENT ON COLUMN fitflow_business_schema.gym_owners.id 
    IS 'Unique owner identifier';

COMMENT ON COLUMN fitflow_business_schema.gym_owners.person_id 
    IS 'Reference to persons table';

COMMENT ON COLUMN fitflow_business_schema.gym_owners.gym_id 
    IS 'Reference to the gym';

COMMENT ON COLUMN fitflow_business_schema.gym_owners.brief_bio 
    IS 'Short description about the owner';

COMMENT ON COLUMN fitflow_business_schema.gym_owners.created_at 
    IS 'Owner creation timestamp';

COMMENT ON COLUMN fitflow_business_schema.gym_owners.updated_at 
    IS 'Last update timestamp';
