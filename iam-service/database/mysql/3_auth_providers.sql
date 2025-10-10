-- Auth Providers Table
-- Stores different authentication providers (google, facebook, apple, mobile_otp, email_otp, local)
CREATE TABLE auth_providers (
    id CHAR(36) PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,
    display_name VARCHAR(100) NOT NULL,
    description TEXT,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert default auth providers
INSERT INTO auth_providers (id, name, display_name, description) VALUES
(UUID(), 'local', 'Local Authentication', 'Username/password authentication'),
(UUID(), 'google', 'Google OAuth', 'Google OAuth 2.0 authentication'),
(UUID(), 'facebook', 'Facebook OAuth', 'Facebook OAuth 2.0 authentication'),
(UUID(), 'apple', 'Apple Sign In', 'Apple Sign In authentication'),
(UUID(), 'mobile_otp', 'Mobile OTP', 'Mobile phone OTP authentication'),
(UUID(), 'email_otp', 'Email OTP', 'Email OTP authentication');

-- Index for performance
CREATE INDEX idx_auth_providers_name ON auth_providers(name);
CREATE INDEX idx_auth_providers_active ON auth_providers(is_active);
