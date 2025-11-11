-- Gym trainers junction table schema for FitFlow Business service
-- This script creates the gym_trainers junction table

-- Create gym_trainers junction table
CREATE TABLE IF NOT EXISTS gym_trainers (
    id BIGSERIAL PRIMARY KEY,
    gym_id BIGINT REFERENCES gyms(id) ON DELETE CASCADE,
    trainer_id BIGINT REFERENCES trainers(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(gym_id, trainer_id)
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_gym_trainers_gym_id ON gym_trainers(gym_id);
CREATE INDEX IF NOT EXISTS idx_gym_trainers_trainer_id ON gym_trainers(trainer_id);
CREATE INDEX IF NOT EXISTS idx_gym_trainers_created_at ON gym_trainers(created_at);

-- Add comments for documentation
COMMENT ON TABLE gym_trainers IS 'Junction table linking gyms and trainers';
COMMENT ON COLUMN gym_trainers.id IS 'Unique relationship identifier';
COMMENT ON COLUMN gym_trainers.gym_id IS 'Reference to the gym';
COMMENT ON COLUMN gym_trainers.trainer_id IS 'Reference to the trainer';
COMMENT ON COLUMN gym_trainers.created_at IS 'Relationship creation timestamp';
