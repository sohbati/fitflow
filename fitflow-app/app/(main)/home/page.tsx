'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { useAuth } from '@/hooks/useAuth'
import { AuthStatus } from '@/components/auth/AuthStatus'
import { GoogleSignInButton } from '@/components/auth/GoogleSignInButton'

export default function HomePage() {
  const router = useRouter()
  const { isAuthenticated, user } = useAuth()
  const [checkingRole, setCheckingRole] = useState(false)

  // Check if user exists in persons table and their role
  useEffect(() => {
    const checkUserStatus = async () => {
      if (!isAuthenticated || !user?.id) return

      setCheckingRole(true)
      try {
        const token = localStorage.getItem('auth_token')
        if (!token) return

        const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'
        
        // First, check if person exists
        const personCheckResponse = await fetch(`${businessServiceUrl}/api/v1/persons/check/${user.id}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (personCheckResponse.ok) {
          const personData = await personCheckResponse.json()
          
          if (!personData.exists) {
            // User doesn't exist in persons table - redirect to role selection
            router.push('/select-role')
            return
          }
        }

        // Person exists, check if they're a gym owner
        const gymOwnerResponse = await fetch(`${businessServiceUrl}/api/v1/gym-owners/user/${user.id}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (gymOwnerResponse.ok) {
          // User is a gym owner, redirect to dashboard
          router.push('/gym-owner/dashboard')
          return
        }
      } catch (error) {
        console.error('Error checking user status:', error)
      } finally {
        setCheckingRole(false)
      }
    }

    checkUserStatus()
  }, [isAuthenticated, user, router])

  if (checkingRole) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div className="text-center">
          <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-600">Loading...</p>
        </div>
      </div>
    )
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

