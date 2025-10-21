'use client'

import { createContext, useContext, ReactNode } from 'react';
import { useGoogleAuth } from '@/hooks/useGoogleAuth';
import { GoogleUser } from '@/lib/google-auth';

interface GoogleAuthContextType {
  user: GoogleUser | null;
  isLoading: boolean;
  error: string | null;
  isGoogleLoaded: boolean;
  signIn: () => void;
  signOut: () => void;
}

const GoogleAuthContext = createContext<GoogleAuthContextType | undefined>(undefined);

export const useGoogleAuthContext = () => {
  const context = useContext(GoogleAuthContext);
  if (context === undefined) {
    throw new Error('useGoogleAuthContext must be used within a GoogleAuthProvider');
  }
  return context;
};

interface GoogleAuthProviderProps {
  children: ReactNode;
}

export const GoogleAuthProvider = ({ children }: GoogleAuthProviderProps) => {
  const googleAuth = useGoogleAuth();

  return (
    <GoogleAuthContext.Provider value={googleAuth}>
      {children}
    </GoogleAuthContext.Provider>
  );
};
