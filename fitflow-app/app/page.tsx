'use client'

import { useState, useEffect } from 'react'
import Link from 'next/link'

export default function Home() {
  const [deferredPrompt, setDeferredPrompt] = useState<any>(null)
  const [showInstallButton, setShowInstallButton] = useState(false)
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState<any>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    // Check authentication status
    const checkAuth = () => {
      const token = localStorage.getItem('auth_token')
      const userData = localStorage.getItem('user_data')
      
      if (token && userData) {
        setIsAuthenticated(true)
        setUser(JSON.parse(userData))
      }
      setIsLoading(false)
    }

    checkAuth()

    const handler = (e: Event) => {
      e.preventDefault()
      setDeferredPrompt(e)
      setShowInstallButton(true)
    }

    window.addEventListener('beforeinstallprompt', handler)

    return () => {
      window.removeEventListener('beforeinstallprompt', handler)
    }
  }, [])

  const handleGoogleSignIn = async () => {
    try {
      // Get Google auth URL from IAM service
      const response = await fetch(`${process.env.NEXT_PUBLIC_IAM_SERVICE_URL}/auth/google/url`)
      if (!response.ok) {
        throw new Error('Failed to get Google auth URL')
      }
      const data = await response.json()
      // Redirect to Google's auth page
      window.location.href = data.url
    } catch (error) {
      console.error('Error initiating Google sign-in:', error)
      alert('Failed to initiate Google sign-in. Please try again.')
    }
  }

  const handleSignOut = () => {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
    setIsAuthenticated(false)
    setUser(null)
  }

  const handleInstallClick = async () => {
    if (!deferredPrompt) return

    deferredPrompt.prompt()
    const { outcome } = await deferredPrompt.userChoice

    if (outcome === 'accepted') {
      console.log('User accepted the install prompt')
    } else {
      console.log('User dismissed the install prompt')
    }

    setDeferredPrompt(null)
    setShowInstallButton(false)
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Navigation */}
      <nav className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <h1 className="text-2xl font-bold text-blue-600">FitFlow</h1>
              </div>
            </div>
            <div className="flex items-center space-x-4">
              <Link 
                href="/" 
                className="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium"
              >
                Home
              </Link>
              <Link 
                href="/about" 
                className="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium"
              >
                About
              </Link>
              {isAuthenticated ? (
                <div className="flex items-center space-x-4">
                  <span className="text-gray-700 text-sm">Welcome, {user?.name || user?.email}</span>
                  <button
                    onClick={handleSignOut}
                    className="bg-red-600 hover:bg-red-700 text-white px-3 py-2 rounded-md text-sm font-medium"
                  >
                    Sign Out
                  </button>
                </div>
              ) : (
                <Link 
                  href="/auth/signin" 
                  className="bg-blue-600 hover:bg-blue-700 text-white px-3 py-2 rounded-md text-sm font-medium"
                >
                  Sign In
                </Link>
              )}
            </div>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="text-center">
          <h1 className="text-4xl font-bold text-gray-900 sm:text-5xl md:text-6xl">
            Welcome to{' '}
            <span className="text-blue-600">FitFlow</span>
          </h1>
          <p className="mt-3 max-w-md mx-auto text-base text-gray-500 sm:text-lg md:mt-5 md:text-xl md:max-w-3xl">
            A modern Progressive Web App for fitness tracking. Track your workouts, 
            monitor your progress, and achieve your fitness goals.
          </p>
          
          <div className="mt-5 max-w-md mx-auto sm:flex sm:justify-center md:mt-8">
            <div className="rounded-md shadow">
              {showInstallButton ? (
                <button
                  onClick={handleInstallClick}
                  className="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 md:py-4 md:text-lg md:px-10"
                >
                  Install App
                </button>
              ) : (
                <div className="w-full flex items-center justify-center px-8 py-3 border border-transparent text-base font-medium rounded-md text-gray-400 bg-gray-200 md:py-4 md:text-lg md:px-10">
                  App Installed
                </div>
              )}
            </div>
          </div>

          {/* Authentication Section */}
          <div className="mt-8 max-w-md mx-auto">
            {isLoading ? (
              <div className="text-center">
                <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto"></div>
                <p className="text-gray-600 text-sm mt-2">Loading...</p>
              </div>
            ) : (
              <>
                <div className="text-center mb-4">
                  <h3 className="text-lg font-semibold text-gray-900">
                    {isAuthenticated ? `Welcome back, ${user?.name || user?.email}!` : 'Sign in to continue'}
                  </h3>
                  <p className="text-gray-600 text-sm mt-1">
                    {isAuthenticated ? 'You are signed in with Google' : 'Sign in with your Google account to access all features'}
                  </p>
                </div>
                
                {isAuthenticated ? (
                  <div className="bg-green-50 border border-green-200 rounded-lg p-4">
                    <div className="flex items-center space-x-3">
                      {user?.picture && (
                        <img
                          src={user.picture}
                          alt={user.name || user.email}
                          className="w-8 h-8 rounded-full"
                        />
                      )}
                      <div className="flex-1">
                        <p className="text-green-800 font-medium">{user?.name || user?.email}</p>
                        <p className="text-green-600 text-sm">{user?.email}</p>
                      </div>
                      <button
                        onClick={handleSignOut}
                        className="text-green-600 hover:text-green-800 text-sm underline"
                      >
                        Sign out
                      </button>
                    </div>
                  </div>
                ) : (
                  <button
                    onClick={handleGoogleSignIn}
                    className="w-full flex items-center justify-center px-4 py-3 border border-gray-300 rounded-lg text-gray-700 bg-white hover:bg-gray-50 transition-colors duration-200"
                  >
                    <div className="flex items-center space-x-2">
                      <svg className="w-5 h-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                      <path d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z" fill="#4285F4"/>
                      <path d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z" fill="#34A853"/>
                      <path d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z" fill="#FBBC05"/>
                      <path d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z" fill="#EA4335"/>
                    </svg>
                      <span>Sign in with Google</span>
                    </div>
                  </button>
                )}
              </>
            )}
          </div>

          {/* Features */}
          <div className="mt-16">
            <h2 className="text-2xl font-bold text-gray-900 mb-8">Features</h2>
            <div className="grid grid-cols-1 gap-8 sm:grid-cols-2 lg:grid-cols-3">
              <div className="bg-white rounded-lg shadow-md p-6">
                <div className="text-blue-600 text-3xl mb-4">📱</div>
                <h3 className="text-lg font-semibold text-gray-900 mb-2">Progressive Web App</h3>
                <p className="text-gray-600">Can be installed on your device like a native app with modern web technologies.</p>
              </div>
              <div className="bg-white rounded-lg shadow-md p-6">
                <div className="text-blue-600 text-3xl mb-4">🏃‍♂️</div>
                <h3 className="text-lg font-semibold text-gray-900 mb-2">Fitness Tracking</h3>
                <p className="text-gray-600">Track your workouts, monitor progress, and set fitness goals.</p>
              </div>
              <div className="bg-white rounded-lg shadow-md p-6">
                <div className="text-blue-600 text-3xl mb-4">📊</div>
                <h3 className="text-lg font-semibold text-gray-900 mb-2">Analytics</h3>
                <p className="text-gray-600">View detailed analytics and insights about your fitness journey.</p>
              </div>
            </div>
          </div>

          {/* Navigation Links */}
          <div className="mt-16">
            <h2 className="text-2xl font-bold text-gray-900 mb-8">Explore</h2>
            <div className="flex flex-col sm:flex-row gap-4 justify-center">
              <Link 
                href="/about"
                className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-blue-600 bg-white hover:bg-gray-50 shadow-md"
              >
                Learn More
              </Link>
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}

