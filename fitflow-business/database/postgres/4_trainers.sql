-- Trainers table schema for FitFlow Business service
-- This script creates the trainers table
-- Each trainer is linked to a person in the persons table

-- Create trainers table
CREATE TABLE IF NOT EXISTS trainers (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    is_registered BOOLEAN DEFAULT FALSE, -- registered with a gym or independent
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_trainers_person_id ON trainers(person_id);
CREATE INDEX IF NOT EXISTS idx_trainers_registered ON trainers(is_registered);

-- Create trigger for updated_at
CREATE TRIGGER update_trainers_updated_at 
    BEFORE UPDATE ON trainers 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE trainers IS 'Trainer information linked to persons table';
COMMENT ON COLUMN trainers.id IS 'Unique trainer identifier';
COMMENT ON COLUMN trainers.person_id IS 'Reference to persons table';
COMMENT ON COLUMN trainers.is_registered IS 'Whether trainer is registered with a gym';
COMMENT ON COLUMN trainers.created_at IS 'Trainer creation timestamp';
COMMENT ON COLUMN trainers.updated_at IS 'Last update timestamp';