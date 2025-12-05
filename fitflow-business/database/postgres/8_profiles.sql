-- Profiles table for FitFlow Business service
-- Allows users to have multiple profiles (gym owner, trainer, trainee)

-- Create schema if it doesn't exist

CREATE TABLE IF NOT EXISTS fitflow_business_schema.profiles (
    id BIGSERIAL PRIMARY KEY,

    user_id UUID NOT NULL, -- Reference to IAM users table
    type VARCHAR(20) NOT NULL, -- 'gym_owner', 'trainer', or 'trainee'
    person_id BIGINT NOT NULL, -- Reference to persons table

    -- Optional references to role-specific entities
    gym_owner_id BIGINT,
    trainer_id BIGINT,
    trainee_id BIGINT,

    is_active BOOLEAN DEFAULT TRUE,
    is_default BOOLEAN DEFAULT FALSE,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_profiles_person 
        FOREIGN KEY (person_id) 
        REFERENCES fitflow_business_schema.persons(id) 
        ON DELETE CASCADE,

    CONSTRAINT fk_profiles_gym_owner 
        FOREIGN KEY (gym_owner_id) 
        REFERENCES fitflow_business_schema.gym_owners(id) 
        ON DELETE SET NULL,

    CONSTRAINT fk_profiles_trainer 
        FOREIGN KEY (trainer_id) 
        REFERENCES fitflow_business_schema.trainers(id) 
        ON DELETE SET NULL,

    CONSTRAINT fk_profiles_trainee 
        FOREIGN KEY (trainee_id) 
        REFERENCES fitflow_business_schema.trainees(id) 
        ON DELETE SET NULL,

    -- Ensure profile type matches the referenced entity
    CONSTRAINT chk_profile_type_match CHECK (
        (type = 'gym_owner' AND gym_owner_id IS NOT NULL) OR
        (type = 'trainer' AND trainer_id IS NOT NULL) OR
        (type = 'trainee' AND trainee_id IS NOT NULL)
    )
);

-- Ensure only one default profile per user (using partial unique index)
CREATE UNIQUE INDEX IF NOT EXISTS idx_profiles_unique_default_per_user 
    ON fitflow_business_schema.profiles(user_id) 
    WHERE is_default = TRUE;

-- Indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_profiles_user_id 
    ON fitflow_business_schema.profiles(user_id);

CREATE INDEX IF NOT EXISTS idx_profiles_type 
    ON fitflow_business_schema.profiles(type);

CREATE INDEX IF NOT EXISTS idx_profiles_person_id 
    ON fitflow_business_schema.profiles(person_id);

CREATE INDEX IF NOT EXISTS idx_profiles_gym_owner_id 
    ON fitflow_business_schema.profiles(gym_owner_id);

CREATE INDEX IF NOT EXISTS idx_profiles_trainer_id 
    ON fitflow_business_schema.profiles(trainer_id);

CREATE INDEX IF NOT EXISTS idx_profiles_trainee_id 
    ON fitflow_business_schema.profiles(trainee_id);

CREATE INDEX IF NOT EXISTS idx_profiles_is_active 
    ON fitflow_business_schema.profiles(is_active);

CREATE INDEX IF NOT EXISTS idx_profiles_is_default 
    ON fitflow_business_schema.profiles(is_default);

CREATE INDEX IF NOT EXISTS idx_profiles_user_type 
    ON fitflow_business_schema.profiles(user_id, type);

-- Composite index for common queries
CREATE INDEX IF NOT EXISTS idx_profiles_user_active 
    ON fitflow_business_schema.profiles(user_id, is_active);

-- Comments
COMMENT ON TABLE fitflow_business_schema.profiles IS 'User profiles allowing multiple roles (gym owner, trainer, trainee)';
COMMENT ON COLUMN fitflow_business_schema.profiles.user_id IS 'Reference to IAM users table';
COMMENT ON COLUMN fitflow_business_schema.profiles.type IS 'Profile type: gym_owner, trainer, or trainee';
COMMENT ON COLUMN fitflow_business_schema.profiles.person_id IS 'Reference to persons table';
COMMENT ON COLUMN fitflow_business_schema.profiles.gym_owner_id IS 'Reference to gym_owners table (for gym_owner type)';
COMMENT ON COLUMN fitflow_business_schema.profiles.trainer_id IS 'Reference to trainers table (for trainer type)';
COMMENT ON COLUMN fitflow_business_schema.profiles.trainee_id IS 'Reference to trainees table (for trainee type)';
COMMENT ON COLUMN fitflow_business_schema.profiles.is_active IS 'Whether the profile is currently active';
COMMENT ON COLUMN fitflow_business_schema.profiles.is_default IS 'Whether this is the default profile for the user';

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION fitflow_business_schema.update_profiles_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger to automatically update updated_at
CREATE TRIGGER trigger_update_profiles_updated_at
    BEFORE UPDATE ON fitflow_business_schema.profiles
    FOR EACH ROW
    EXECUTE FUNCTION fitflow_business_schema.update_profiles_updated_at();

