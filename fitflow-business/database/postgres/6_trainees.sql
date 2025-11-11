-- Trainee table for FitFlow Business service
-- This table stores gym-specific information for trainees/members
-- Each trainee is linked to a person in the persons table

-- Create the trainee table
CREATE TABLE IF NOT EXISTS trainees (
    id BIGSERIAL PRIMARY KEY,
    person_id BIGINT NOT NULL REFERENCES persons(id) ON DELETE CASCADE,
    height_cm INTEGER, -- height in centimeters
    weight_kg DECIMAL(5,2), -- weight in kilograms
    fitness_level VARCHAR(20) DEFAULT 'beginner', -- beginner, intermediate, advanced
    goals TEXT, -- fitness goals and objectives
    medical_conditions TEXT, -- any medical conditions or restrictions
    membership_type VARCHAR(50) DEFAULT 'basic', -- basic, premium, vip
    membership_start_date DATE,
    membership_end_date DATE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_trainees_person_id ON trainees(person_id);
CREATE INDEX IF NOT EXISTS idx_trainees_membership_type ON trainees(membership_type);
CREATE INDEX IF NOT EXISTS idx_trainees_is_active ON trainees(is_active);
CREATE INDEX IF NOT EXISTS idx_trainees_membership_dates ON trainees(membership_start_date, membership_end_date);
CREATE INDEX IF NOT EXISTS idx_trainees_fitness_level ON trainees(fitness_level);

-- Create trigger for updated_at timestamp
CREATE TRIGGER update_trainees_updated_at 
    BEFORE UPDATE ON trainees 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE trainees IS 'Stores gym-specific information for trainees/members, linked to persons table';
COMMENT ON COLUMN trainees.person_id IS 'Reference to persons table';
COMMENT ON COLUMN trainees.height_cm IS 'Height in centimeters';
COMMENT ON COLUMN trainees.weight_kg IS 'Weight in kilograms';
COMMENT ON COLUMN trainees.fitness_level IS 'Current fitness level: beginner, intermediate, advanced';
COMMENT ON COLUMN trainees.goals IS 'Fitness goals and objectives';
COMMENT ON COLUMN trainees.medical_conditions IS 'Any medical conditions or restrictions';
COMMENT ON COLUMN trainees.membership_type IS 'Type of membership: basic, premium, vip';
COMMENT ON COLUMN trainees.membership_start_date IS 'Membership start date';
COMMENT ON COLUMN trainees.membership_end_date IS 'Membership end date';
COMMENT ON COLUMN trainees.is_active IS 'Whether the trainee is currently active';
