-- Gym locations table schema for FitFlow Business service
-- This script creates the gym_locations table

-- Create gym_locations table
CREATE TABLE IF NOT EXISTS gym_locations (
    id BIGSERIAL PRIMARY KEY,
    gym_id BIGINT REFERENCES gyms(id) ON DELETE CASCADE,
    location_type VARCHAR(20) NOT NULL,  -- home, gym, park, nature
    address VARCHAR(255),
    city VARCHAR(100),
    province VARCHAR(100),
    country VARCHAR(100),
    postal_code VARCHAR(20),
    latitude DECIMAL(10,7),
    longitude DECIMAL(10,7),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for better performance
CREATE INDEX IF NOT EXISTS idx_gym_locations_gym_id ON gym_locations(gym_id);
CREATE INDEX IF NOT EXISTS idx_gym_locations_type ON gym_locations(location_type);
CREATE INDEX IF NOT EXISTS idx_gym_locations_city ON gym_locations(city);
CREATE INDEX IF NOT EXISTS idx_gym_locations_coordinates ON gym_locations(latitude, longitude);

-- Create trigger for updated_at
CREATE TRIGGER update_gym_locations_updated_at 
    BEFORE UPDATE ON gym_locations 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();

-- Add comments for documentation
COMMENT ON TABLE gym_locations IS 'Gym location details including addresses and coordinates';
COMMENT ON COLUMN gym_locations.id IS 'Unique location identifier';
COMMENT ON COLUMN gym_locations.gym_id IS 'Reference to the gym';
COMMENT ON COLUMN gym_locations.location_type IS 'Type of location (home, gym, park, nature)';
COMMENT ON COLUMN gym_locations.address IS 'Full address';
COMMENT ON COLUMN gym_locations.city IS 'City name';
COMMENT ON COLUMN gym_locations.province IS 'Province/state name';
COMMENT ON COLUMN gym_locations.country IS 'Country name';
COMMENT ON COLUMN gym_locations.postal_code IS 'Postal/ZIP code';
COMMENT ON COLUMN gym_locations.latitude IS 'Latitude coordinate';
COMMENT ON COLUMN gym_locations.longitude IS 'Longitude coordinate';
COMMENT ON COLUMN gym_locations.created_at IS 'Location creation timestamp';
COMMENT ON COLUMN gym_locations.updated_at IS 'Last update timestamp';
