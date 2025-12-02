# How to Find Your Google OAuth Client ID and Secret

## Step-by-Step Navigation Guide

### Step 1: Sign in to Google Cloud Console
1. Go to: https://console.cloud.google.com/
2. Sign in with your Google account (sohbati@gmail.com)

### Step 2: Select Your Project
1. At the top of the page, click the **project dropdown** (shows current project name)
2. Select the project where you created the OAuth credentials

### Step 3: Navigate to APIs & Services > Credentials
**Option A: Using the Navigation Menu (Hamburger Menu)**
1. Click the **☰ (hamburger menu)** icon in the top-left corner
2. Hover over or click **"APIs & Services"**
3. Click **"Credentials"** in the submenu

**Option B: Direct URL**
- Go directly to: https://console.cloud.google.com/apis/credentials

### Step 4: Find Your OAuth 2.0 Client
1. In the **"Credentials"** page, you'll see a list of credentials
2. Look for **"OAuth 2.0 Client IDs"** section
3. Find your client (it might be named "FitFlow Web Client" or similar)
4. Click on the **pencil/edit icon** (✏️) or click on the client name

### Step 5: Copy Your Credentials
In the OAuth client details page, you'll see:

1. **Client ID**: 
   - Format: `xxxxx-xxxxx.apps.googleusercontent.com`
   - Click the **copy icon** (📋) next to it, or select and copy manually

2. **Client secret**: 
   - Click **"Show"** or **"Reveal"** button to see it
   - Format: `GOCSPX-xxxxxxxxxxxxx`
   - Click the **copy icon** (📋) next to it, or select and copy manually

### Step 6: Update Your Configuration

**For IAM Service** (`iam-service/.env`):
```bash
GOOGLE_CLIENT_ID=paste-your-client-id-here.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=paste-your-client-secret-here
GOOGLE_REDIRECT_URL=http://localhost:3000/auth/google/callback
```

**Important Notes:**
- Make sure there are **no spaces** before or after the `=` sign
- Don't include quotes around the values
- The Client ID ends with `.apps.googleusercontent.com`
- The Client Secret starts with `GOCSPX-`

### Step 7: Verify Redirect URI
While you're in the OAuth client settings, check:
- **Authorized redirect URIs** should include:
  - `http://localhost:3000/auth/google/callback`
  - (Optional) `http://192.168.100.64:3000/auth/google/callback` if accessing from network

### Step 8: Restart IAM Service
After updating `.env`, restart the IAM service:
```bash
# Stop current service (Ctrl+C)
./1-run-iam-service.sh
```

## Quick Visual Guide

```
Google Cloud Console
  └─ ☰ Menu (top-left)
     └─ APIs & Services
        └─ Credentials
           └─ OAuth 2.0 Client IDs
              └─ [Your Client Name]
                 ├─ Client ID: [Copy this]
                 └─ Client secret: [Show] → [Copy this]
```

## Troubleshooting

**Can't find Credentials page?**
- Make sure you're in the correct project
- Check that you have the right permissions (Editor or Owner role)

**Don't see your OAuth client?**
- Check if you're in the right project
- Look in the "OAuth 2.0 Client IDs" section (not API keys or service accounts)

**Client secret shows as hidden?**
- Click "Show" or "Reveal" button
- If you can't see it, you may need to create a new client or regenerate the secret

