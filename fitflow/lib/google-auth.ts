// Google OAuth Configuration
export const GOOGLE_CLIENT_ID = process.env.NEXT_PUBLIC_GOOGLE_CLIENT_ID || '665428305045-jdvt22tb51mrf588vov4523eici6s6a3.apps.googleusercontent.com';

// Google OAuth Scopes
export const GOOGLE_SCOPES = [
  'https://www.googleapis.com/auth/userinfo.email',
  'https://www.googleapis.com/auth/userinfo.profile',
  'openid'
].join(' ');

// Google OAuth Configuration
export const GOOGLE_AUTH_CONFIG = {
  client_id: GOOGLE_CLIENT_ID,
  scope: GOOGLE_SCOPES,
  redirect_uri: typeof window !== 'undefined' ? window.location.origin : '',
  response_type: 'code',
  access_type: 'offline',
  prompt: 'consent'
};

// User profile interface
export interface GoogleUser {
  id: string;
  email: string;
  name: string;
  picture: string;
  given_name?: string;
  family_name?: string;
}

// Auth response interface
export interface AuthResponse {
  user: GoogleUser;
  token: string;
  expires_in: number;
}
