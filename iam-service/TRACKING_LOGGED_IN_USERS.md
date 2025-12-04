# Tracking Logged-In Users in IAM Service

## Current Data Model Limitation

**The current IAM service data model does NOT directly support tracking logged-in users** because:

1. **Stateless JWT Tokens**: The system uses stateless JWT tokens that are:
   - Not stored in the database
   - Validated by signature and expiration time only
   - Default expiration: 15 minutes (configurable via `TOKEN_EXP_MINUTES`)

2. **No Session Storage**: There's no `sessions` table or similar mechanism to track active logins

3. **No Last Login Tracking**: The `users` table doesn't have a `last_login_at` field

## Current Tables

### `users` table
- `id`, `email`, `mobile`, `google_id`, `display_name`, `avatar_url`, `country`, `role`, `is_active`
- `created_at`, `updated_at`
- **No login tracking fields**

### `user_auth` table
- Links users to authentication providers (Google, Local, etc.)
- Stores provider-specific data (passwords, OTP codes)
- **No session or login tracking**

## Solutions to Track Logged-In Users

### Option 1: Add Sessions Table (Recommended)

Create a `sessions` table to track active sessions:

```sql
-- See: database/postgres/5_sessions.sql
```

**Pros:**
- Accurate count of active users
- Can track device, IP, last activity
- Can implement logout functionality
- Can revoke sessions

**Cons:**
- Requires storing token hashes
- Adds database overhead
- Requires cleanup of expired sessions

**Query to count active users:**
```sql
SELECT COUNT(DISTINCT user_id) 
FROM sessions 
WHERE expires_at > CURRENT_TIMESTAMP;
```

### Option 2: Add Last Login Timestamp

Add a `last_login_at` field to the `users` table:

```sql
ALTER TABLE users ADD COLUMN last_login_at TIMESTAMP;
CREATE INDEX idx_users_last_login ON users(last_login_at);
```

**Pros:**
- Simple implementation
- Tracks when users last logged in
- Can estimate active users (users who logged in within token expiration time)

**Cons:**
- Not accurate (doesn't track if user is still active)
- Doesn't account for multiple devices/sessions
- Can't count exact number of logged-in users

**Query to estimate active users:**
```sql
SELECT COUNT(*) 
FROM users 
WHERE last_login_at > (CURRENT_TIMESTAMP - INTERVAL '15 minutes')
  AND is_active = TRUE;
```

### Option 3: Token Blacklist/Whitelist Table

Store active tokens in a database table:

```sql
CREATE TABLE active_tokens (
    token_hash VARCHAR(255) PRIMARY KEY,
    user_id UUID NOT NULL,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
```

**Pros:**
- Can track exact active tokens
- Can implement token revocation
- Accurate count

**Cons:**
- Defeats the purpose of stateless JWTs
- High database overhead
- Requires cleanup

### Option 4: Application-Level Tracking (Redis/Memory)

Use an in-memory store (Redis) or application memory to track active sessions:

**Pros:**
- Fast lookups
- No database overhead
- Can set TTL automatically

**Cons:**
- Lost on server restart (unless using Redis)
- Not persistent across multiple servers (unless using shared Redis)

## Recommended Approach

**For production use, I recommend Option 1 (Sessions Table)** because:

1. **Accurate tracking**: Know exactly who is logged in
2. **Security**: Can revoke sessions, track suspicious activity
3. **Analytics**: Track user activity, device usage
4. **Multi-device support**: Track multiple sessions per user
5. **Logout functionality**: Properly implement logout

## Implementation Steps

1. **Add sessions table** (see `5_sessions.sql`)
2. **Update auth handlers** to create sessions on login
3. **Update JWT middleware** to update `last_activity_at` on each request
4. **Add cleanup job** to remove expired sessions
5. **Add API endpoint** to query active sessions

## Example Queries

### Count Active Logged-In Users
```sql
SELECT COUNT(DISTINCT user_id) as active_users
FROM sessions
WHERE expires_at > CURRENT_TIMESTAMP;
```

### List Active Users with Details
```sql
SELECT 
    u.id,
    u.email,
    u.display_name,
    COUNT(s.id) as active_sessions,
    MAX(s.last_activity_at) as last_activity
FROM users u
INNER JOIN sessions s ON u.id = s.user_id
WHERE s.expires_at > CURRENT_TIMESTAMP
GROUP BY u.id, u.email, u.display_name
ORDER BY last_activity DESC;
```

### Count Sessions by Device
```sql
SELECT 
    device_info,
    COUNT(*) as session_count
FROM sessions
WHERE expires_at > CURRENT_TIMESTAMP
GROUP BY device_info;
```

## Current Workaround

**Without modifying the database**, you can only estimate active users by:

1. Checking JWT token expiration times (requires storing tokens somewhere)
2. Using application logs to track login events
3. Implementing a separate tracking service

**Note**: The current system cannot accurately count logged-in users without adding session tracking.

