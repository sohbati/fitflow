-- User Authentication Table
-- Links users to their authentication providers and stores provider-specific data
CREATE TABLE user_auth (
    id CHAR(36) PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    auth_provider_id CHAR(36) NOT NULL,
    
    -- Provider-specific identifiers
    provider_user_id VARCHAR(255), -- Google ID, Facebook ID, etc.
    provider_email VARCHAR(255),   -- Email from provider (may differ from user email)
    
    -- Provider-specific data
    provider_data JSON,           -- Store additional provider-specific data
    
    -- Verification status
    is_verified BOOLEAN DEFAULT FALSE,
    verified_at TIMESTAMP NULL,
    
    -- OTP fields (for mobile/email OTP)
    otp_code VARCHAR(10) NULL,
    otp_expires_at TIMESTAMP NULL,
    otp_attempts INT DEFAULT 0,
    
    -- Timestamps
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (auth_provider_id) REFERENCES auth_providers(id) ON DELETE CASCADE,
    
    -- Unique constraints
    UNIQUE KEY unique_user_provider (user_id, auth_provider_id),
    UNIQUE KEY unique_provider_user_id (auth_provider_id, provider_user_id)
);

-- Indexes for performance
CREATE INDEX idx_user_auth_user_id ON user_auth(user_id);
CREATE INDEX idx_user_auth_provider_id ON user_auth(auth_provider_id);
CREATE INDEX idx_user_auth_provider_user_id ON user_auth(provider_user_id);
CREATE INDEX idx_user_auth_email ON user_auth(provider_email);
CREATE INDEX idx_user_auth_verified ON user_auth(is_verified);
CREATE INDEX idx_user_auth_otp ON user_auth(otp_code, otp_expires_at);

-- Add username field for local authentication
ALTER TABLE user_auth ADD COLUMN username VARCHAR(255) NULL;
CREATE INDEX idx_user_auth_username ON user_auth(username);
