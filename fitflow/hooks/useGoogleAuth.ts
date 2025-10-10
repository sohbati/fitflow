'use client'

import { useState, useEffect, useCallback } from 'react';
import { GoogleUser, AuthResponse, GOOGLE_CLIENT_ID } from '@/lib/google-auth';

declare global {
  interface Window {
    google: any;
    gapi: any;
  }
}

export const useGoogleAuth = () => {
  const [user, setUser] = useState<GoogleUser | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isGoogleLoaded, setIsGoogleLoaded] = useState(false);

  // Load Google Identity Services
  useEffect(() => {
    const loadGoogleScript = () => {
      if (window.google) {
        setIsGoogleLoaded(true);
        return;
      }

      const script = document.createElement('script');
      script.src = 'https://accounts.google.com/gsi/client';
      script.async = true;
      script.defer = true;
      script.onload = () => {
        setIsGoogleLoaded(true);
      };
      script.onerror = () => {
        setError('Failed to load Google Identity Services');
      };
      document.head.appendChild(script);
    };

    loadGoogleScript();
  }, []);

  // Initialize Google Identity Services
  useEffect(() => {
    if (!isGoogleLoaded || !window.google) return;

    try {
      window.google.accounts.id.initialize({
        client_id: GOOGLE_CLIENT_ID,
        callback: handleCredentialResponse,
        auto_select: false,
        cancel_on_tap_outside: true,
      });
    } catch (err) {
      setError('Failed to initialize Google Identity Services');
    }
  }, [isGoogleLoaded]);

  // Handle credential response
  const handleCredentialResponse = useCallback((response: any) => {
    setIsLoading(true);
    setError(null);

    try {
      // Decode the JWT token to get user info
      const payload = JSON.parse(atob(response.credential.split('.')[1]));
      
      const userData: GoogleUser = {
        id: payload.sub,
        email: payload.email,
        name: payload.name,
        picture: payload.picture,
        given_name: payload.given_name,
        family_name: payload.family_name,
      };

      setUser(userData);
      
      // Store token in localStorage
      localStorage.setItem('google_token', response.credential);
      localStorage.setItem('google_user', JSON.stringify(userData));
      
    } catch (err) {
      setError('Failed to process Google authentication');
    } finally {
      setIsLoading(false);
    }
  }, []);

  // Sign in with Google
  const signIn = useCallback(() => {
    if (!isGoogleLoaded || !window.google) {
      setError('Google Identity Services not loaded');
      return;
    }

    setIsLoading(true);
    setError(null);

    try {
      window.google.accounts.id.prompt();
    } catch (err) {
      setError('Failed to initiate Google sign-in');
      setIsLoading(false);
    }
  }, [isGoogleLoaded]);

  // Sign out
  const signOut = useCallback(() => {
    if (!isGoogleLoaded || !window.google) return;

    try {
      window.google.accounts.id.disableAutoSelect();
      setUser(null);
      localStorage.removeItem('google_token');
      localStorage.removeItem('google_user');
    } catch (err) {
      setError('Failed to sign out');
    }
  }, [isGoogleLoaded]);

  // Check for existing session
  useEffect(() => {
    const storedUser = localStorage.getItem('google_user');
    const storedToken = localStorage.getItem('google_token');

    if (storedUser && storedToken) {
      try {
        const userData = JSON.parse(storedUser);
        setUser(userData);
      } catch (err) {
        // Clear invalid stored data
        localStorage.removeItem('google_user');
        localStorage.removeItem('google_token');
      }
    }
  }, []);

  return {
    user,
    isLoading,
    error,
    isGoogleLoaded,
    signIn,
    signOut,
  };
};
