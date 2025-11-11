-- Gym owners table schema for FitFlow Business service
-- This script creates the gym_owners table
-- Each gym owner is linked to a person in the persons table

-- Create gym_owners table
CREATE TABLE IF NOT EXISTS gym_owners (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    gym_id BIGINT NOT NULL REFERENCES gyms(id) ON DELETE CASCADE,
    brief_bio TEXT,           -- short description about the owner
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_gym_owners_person_id ON gym_owners(person_id);
CREATE INDEX IF NOT EXISTS idx_gym_owners_gym_id ON gym_owners(gym_id);
CREATE INDEX IF NOT EXISTS idx_gym_owners_person_gym ON gym_owners(person_id, gym_id);

-- Create trigger for updated_at
CREATE TRIGGER update_gym_owners_updated_at 
    BEFORE UPDATE ON gym_owners 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE gym_owners IS 'Gym owner information linked to persons table';
COMMENT ON COLUMN gym_owners.id IS 'Unique owner identifier';
COMMENT ON COLUMN gym_owners.person_id IS 'Reference to persons table';
COMMENT ON COLUMN gym_owners.gym_id IS 'Reference to the gym';
COMMENT ON COLUMN gym_owners.brief_bio IS 'Short description about the owner';
COMMENT ON COLUMN gym_owners.created_at IS 'Owner creation timestamp';
COMMENT ON COLUMN gym_owners.updated_at IS 'Last update timestamp';