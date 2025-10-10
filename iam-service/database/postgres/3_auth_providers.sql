-- Auth Providers Table
-- Stores different authentication providers (google, facebook, apple, mobile_otp, email_otp, local)
CREATE TABLE auth_providers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(50) NOT NULL UNIQUE,
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert default auth providers
INSERT INTO auth_providers (name, display_name, description) VALUES
('local', 'Local Authentication', 'Username/password authentication'),
('google', 'Google OAuth', 'Google OAuth 2.0 authentication'),
('facebook', 'Facebook OAuth', 'Facebook OAuth 2.0 authentication'),
('apple', 'Apple Sign In', 'Apple Sign In authentication'),
('mobile_otp', 'Mobile OTP', 'Mobile phone OTP authentication'),
('email_otp', 'Email OTP', 'Email OTP authentication');

-- Index for performance
CREATE INDEX idx_auth_providers_name ON auth_providers(name);
CREATE INDEX idx_auth_providers_active ON auth_providers(is_active);

-- Update trigger for updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_auth_providers_updated_at BEFORE UPDATE ON auth_providers FOR EACH ROW EXECUTE FUNCTION update_updated_at_column();
