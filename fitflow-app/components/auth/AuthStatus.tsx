'use client'

import { useAuth } from '@/hooks/useAuth'
import { GoogleSignInButton } from './GoogleSignInButton'

export function AuthStatus() {
  const { isAuthenticated, user, isLoading, signOut } = useAuth()

  if (isLoading) {
    return (
      <div className="text-center">
        <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto"></div>
        <p className="text-gray-600 text-sm mt-2">Loading...</p>
      </div>
    )
  }

  if (isAuthenticated) {
    return (
      <div className="bg-green-50 border border-green-200 rounded-lg p-4">
        <div className="flex items-center space-x-3">
          {user?.picture && (
            <img
              src={user.picture}
              alt={user.name || user.email || 'User'}
              className="w-8 h-8 rounded-full"
            />
          )}
          <div className="flex-1">
            <p className="text-green-800 font-medium">{user?.name || user?.email}</p>
            <p className="text-green-600 text-sm">{user?.email}</p>
          </div>
          <button
            onClick={signOut}
            className="text-green-600 hover:text-green-800 text-sm underline"
          >
            Sign out
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="text-center space-y-4">
      <div>
        <h3 className="text-lg font-semibold text-gray-900">Sign in to continue</h3>
        <p className="text-gray-600 text-sm mt-1">
          Sign in with your Google account to access all features
        </p>
      </div>
      <GoogleSignInButton fullWidth />
    </div>
  )
}

