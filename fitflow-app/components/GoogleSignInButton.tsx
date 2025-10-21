'use client'

import { useGoogleAuth } from '@/hooks/useGoogleAuth';

interface GoogleSignInButtonProps {
  onSuccess?: (user: any) => void;
  onError?: (error: string) => void;
  className?: string;
  text?: string;
}

export const GoogleSignInButton = ({ 
  onSuccess, 
  onError, 
  className = '',
  text = 'Sign in with Google'
}: GoogleSignInButtonProps) => {
  const { user, isLoading, error, isGoogleLoaded, signIn, signOut } = useGoogleAuth();

  // Handle success callback
  const handleSuccess = (userData: any) => {
    if (onSuccess) {
      onSuccess(userData);
    }
  };

  // Handle error callback
  const handleError = (errorMessage: string) => {
    if (onError) {
      onError(errorMessage);
    }
  };

  // Handle sign in click
  const handleSignIn = () => {
    if (!isGoogleLoaded) {
      handleError('Google Identity Services not loaded');
      return;
    }
    signIn();
  };

  // Handle sign out click
  const handleSignOut = () => {
    signOut();
  };

  // Show error if any
  if (error) {
    return (
      <div className={`bg-red-50 border border-red-200 rounded-lg p-4 ${className}`}>
        <p className="text-red-800 text-sm">{error}</p>
        <button
          onClick={handleSignIn}
          className="mt-2 text-red-600 hover:text-red-800 text-sm underline"
        >
          Try again
        </button>
      </div>
    );
  }

  // Show user info if signed in
  if (user) {
    return (
      <div className={`bg-green-50 border border-green-200 rounded-lg p-4 ${className}`}>
        <div className="flex items-center space-x-3">
          <img
            src={user.picture}
            alt={user.name}
            className="w-8 h-8 rounded-full"
          />
          <div className="flex-1">
            <p className="text-green-800 font-medium">{user.name}</p>
            <p className="text-green-600 text-sm">{user.email}</p>
          </div>
          <button
            onClick={handleSignOut}
            className="text-green-600 hover:text-green-800 text-sm underline"
          >
            Sign out
          </button>
        </div>
      </div>
    );
  }

  // Show sign in button
  return (
    <button
      onClick={handleSignIn}
      disabled={isLoading || !isGoogleLoaded}
      className={`
        w-full flex items-center justify-center px-4 py-3 border border-gray-300 
        rounded-lg text-gray-700 bg-white hover:bg-gray-50 
        disabled:opacity-50 disabled:cursor-not-allowed
        transition-colors duration-200
        ${className}
      `}
    >
      {isLoading ? (
        <div className="flex items-center space-x-2">
          <div className="w-4 h-4 border-2 border-gray-300 border-t-blue-600 rounded-full animate-spin"></div>
          <span>Signing in...</span>
        </div>
      ) : (
        <div className="flex items-center space-x-2">
          <svg
            className="w-5 h-5"
            viewBox="0 0 24 24"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"
              fill="#4285F4"
            />
            <path
              d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"
              fill="#34A853"
            />
            <path
              d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"
              fill="#FBBC05"
            />
            <path
              d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"
              fill="#EA4335"
            />
          </svg>
          <span>{text}</span>
        </div>
      )}
    </button>
  );
};
