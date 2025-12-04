-- Sessions Table
-- Tracks active user sessions for login monitoring
-- This table is optional and can be added if you need to track active logins

CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL,
    token_hash VARCHAR(255) NOT NULL, -- Hash of the JWT token (for security)
    device_info VARCHAR(500),          -- Device/browser information
    ip_address VARCHAR(45),            -- IPv4 or IPv6 address
    user_agent TEXT,                    -- Browser user agent
    expires_at TIMESTAMP NOT NULL,      -- When the session expires
    last_activity_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Last activity timestamp
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraint
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    
    -- Indexes
    UNIQUE (token_hash)
);

-- Indexes for performance
CREATE INDEX IF NOT EXISTS idx_sessions_user_id ON sessions(user_id);
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);
CREATE INDEX IF NOT EXISTS idx_sessions_last_activity ON sessions(last_activity_at);
-- Composite index for efficient queries on active sessions (user_id, expires_at)
-- Note: We can't use CURRENT_TIMESTAMP in WHERE clause (not immutable), but this index
-- will still be efficient for queries filtering by expires_at > CURRENT_TIMESTAMP
CREATE INDEX IF NOT EXISTS idx_sessions_active ON sessions(user_id, expires_at);

-- Add comments
COMMENT ON TABLE sessions IS 'Active user sessions for tracking logged-in users';
COMMENT ON COLUMN sessions.token_hash IS 'SHA-256 hash of the JWT token';
COMMENT ON COLUMN sessions.expires_at IS 'Session expiration time (matches JWT expiration)';
COMMENT ON COLUMN sessions.last_activity_at IS 'Last time the user made a request with this session';
 
