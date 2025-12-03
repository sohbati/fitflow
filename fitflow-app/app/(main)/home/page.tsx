'use client'

import { useState, useEffect } from 'react'
import Link from 'next/link'
import { useAuth } from '@/hooks/useAuth'
import { AuthStatus } from '@/components/auth/AuthStatus'
import { GoogleSignInButton } from '@/components/auth/GoogleSignInButton'

export default function HomePage() {
  const { isAuthenticated, user } = useAuth()
  const [deferredPrompt, setDeferredPrompt] = useState<any>(null)
  const [showInstallButton, setShowInstallButton] = useState(false)

  useEffect(() => {
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
        <AuthStatus />
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
  )
}

