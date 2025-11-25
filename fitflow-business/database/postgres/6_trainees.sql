-- Create trainees table inside schema
CREATE TABLE IF NOT EXISTS fitflow_business_schema.trainees (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL ,
    height_cm INTEGER,                          -- height in centimeters
    weight_kg DECIMAL(5,2),                     -- weight in kilograms
    fitness_level VARCHAR(20) DEFAULT 'beginner',  -- beginner, intermediate, advanced
    goals TEXT,                                 -- fitness goals and objectives
    medical_conditions TEXT,                    -- medical conditions or restrictions
    membership_type VARCHAR(50) DEFAULT 'basic',   -- basic, premium, vip
    membership_start_date DATE,
    membership_end_date DATE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_trainees_person_id 
    ON fitflow_business_schema.trainees(person_id);

CREATE INDEX IF NOT EXISTS idx_trainees_membership_type 
    ON fitflow_business_schema.trainees(membership_type);

CREATE INDEX IF NOT EXISTS idx_trainees_is_active 
    ON fitflow_business_schema.trainees(is_active);

CREATE INDEX IF NOT EXISTS idx_trainees_membership_dates 
    ON fitflow_business_schema.trainees(membership_start_date, membership_end_date);

CREATE INDEX IF NOT EXISTS idx_trainees_fitness_level 
    ON fitflow_business_schema.trainees(fitness_level);

-- Add comments for documentation
COMMENT ON TABLE fitflow_business_schema.trainees 
    IS 'Stores gym-specific information for trainees/members, linked to persons table';

COMMENT ON COLUMN fitflow_business_schema.trainees.person_id 
    IS 'Reference to persons table';

COMMENT ON COLUMN fitflow_business_schema.trainees.height_cm 
    IS 'Height in centimeters';

COMMENT ON COLUMN fitflow_business_schema.trainees.weight_kg 
    IS 'Weight in kilograms';

COMMENT ON COLUMN fitflow_business_schema.trainees.fitness_level 
    IS 'Current fitness level: beginner, intermediate, advanced';

COMMENT ON COLUMN fitflow_business_schema.trainees.goals 
    IS 'Fitness goals and objectives';

COMMENT ON COLUMN fitflow_business_schema.trainees.medical_conditions 
    IS 'Any medical conditions or restrictions';

COMMENT ON COLUMN fitflow_business_schema.trainees.membership_type 
    IS 'Type of membership: basic, premium, vip';

COMMENT ON COLUMN fitflow_business_schema.trainees.membership_start_date 
    IS 'Membership start date';

COMMENT ON COLUMN fitflow_business_schema.trainees.membership_end_date 
    IS 'Membership end date';

COMMENT ON COLUMN fitflow_business_schema.trainees.is_active 
    IS 'Whether the trainee is currently active';
