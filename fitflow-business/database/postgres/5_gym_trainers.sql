-- Create gym_trainers junction table inside schema
CREATE TABLE IF NOT EXISTS fitflow_business_schema.gym_trainers (
    id BIGSERIAL PRIMARY KEY,
    gym_id BIGINT 
        REFERENCES fitflow_business_schema.gyms(id) ON DELETE CASCADE,
    trainer_id BIGINT 
        REFERENCES fitflow_business_schema.trainers(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(gym_id, trainer_id)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_gym_trainers_gym_id 
    ON fitflow_business_schema.gym_trainers(gym_id);

CREATE INDEX IF NOT EXISTS idx_gym_trainers_trainer_id 
    ON fitflow_business_schema.gym_trainers(trainer_id);

CREATE INDEX IF NOT EXISTS idx_gym_trainers_created_at 
    ON fitflow_business_schema.gym_trainers(created_at);

-- Add comments for documentation
COMMENT ON TABLE fitflow_business_schema.gym_trainers 
    IS 'Junction table linking gyms and trainers';

COMMENT ON COLUMN fitflow_business_schema.gym_trainers.id 
    IS 'Unique relationship identifier';

COMMENT ON COLUMN fitflow_business_schema.gym_trainers.gym_id 
    IS 'Reference to the gym';

COMMENT ON COLUMN fitflow_business_schema.gym_trainers.trainer_id 
    IS 'Reference to the trainer';

COMMENT ON COLUMN fitflow_business_schema.gym_trainers.created_at 
    IS 'Relationship creation timestamp';
