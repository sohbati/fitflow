# Google OAuth Setup Guide

## Step 1: Create Google OAuth Credentials

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Enable the Google+ API:
   - Go to "APIs & Services" > "Library"
   - Search for "Google+ API" or "People API"
   - Click "Enable"

4. Create OAuth 2.0 Credentials:
   - Go to "APIs & Services" > "Credentials"
   - Click "Create Credentials" > "OAuth client ID"
   - If prompted, configure the OAuth consent screen first:
     - User Type: External (for testing) or Internal (for organization)
     - App name: FitFlow
     - User support email: your-email@gmail.com
     - Developer contact: your-email@gmail.com
     - Click "Save and Continue"
     - Scopes: Add "userinfo.email" and "userinfo.profile"
     - Test users: Add your email address
     - Click "Save and Continue"

5. Create OAuth Client:
   - Application type: "Web application"
   - Name: FitFlow Web Client
   - Authorized JavaScript origins:
     - `http://localhost:3000`
     - `http://192.168.100.64:3000` (if accessing from network)
   - Authorized redirect URIs:
     - `http://localhost:3000/auth/google/callback`
     - `http://192.168.100.64:3000/auth/google/callback` (if accessing from network)
   - Click "Create"
   - Copy the **Client ID** and **Client Secret**

## Step 2: Update Configuration Files

### Update IAM Service Configuration

Edit `iam-service/.env`:

```bash
GOOGLE_CLIENT_ID=your-actual-client-id-here.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=your-actual-client-secret-here
GOOGLE_REDIRECT_URL=http://localhost:3000/auth/google/callback
```

### Update Next.js App Configuration (Optional)

If you need Google OAuth in the Next.js app directly, edit `fitflow-app/.env.local`:

```bash
GOOGLE_CLIENT_ID=your-actual-client-id-here.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=your-actual-client-secret-here
```

## Step 3: Restart Services

After updating the configuration:

1. **Restart IAM Service:**
   ```bash
   # Stop the current IAM service (Ctrl+C)
   ./1-run-iam-service.sh
   ```

2. **Restart Next.js App (if you updated .env.local):**
   ```bash
   # Stop the current dev server (Ctrl+C)
   cd fitflow-app
   npm run dev
   ```

## Step 4: Test

1. Go to your app: http://localhost:3000
2. Click "Sign in with Google"
3. You should be redirected to Google's sign-in page
4. After signing in, you'll be redirected back to your app

## Troubleshooting

- **Error 401: invalid_client**: Check that Client ID and Secret are correct
- **Redirect URI mismatch**: Ensure the redirect URI in Google Console matches exactly: `http://localhost:3000/auth/google/callback`
- **CORS errors**: Already fixed in the gateway configuration

