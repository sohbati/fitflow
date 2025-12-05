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

  // Check user profiles and redirect accordingly
  useEffect(() => {
    const checkUserProfiles = async () => {
      if (!isAuthenticated || !user?.id) return

      setCheckingRole(true)
      try {
        const token = localStorage.getItem('auth_token')
        if (!token) return

        const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'
        
        // Check for selected profile first (only if it exists and is valid)
        const selectedProfileId = localStorage.getItem('selected_profile_id')
        
        if (selectedProfileId) {
          // User has a selected profile, verify it's still valid
          const profileResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/${selectedProfileId}`, {
            headers: {
              'Authorization': `Bearer ${token}`,
            },
          })

          if (profileResponse.ok) {
            const profile = await profileResponse.json()
            // Verify the profile belongs to this user and is active
            if (profile.user_id === user.id && profile.is_active) {
              switch (profile.type) {
                case 'gym_owner':
                  router.push('/gym-owner')
                  return
                case 'trainer':
                  router.push('/trainer')
                  return
                case 'trainee':
                  router.push('/trainee')
                  return
              }
            } else {
              // Profile is invalid, clear it
              localStorage.removeItem('selected_profile_id')
            }
          } else {
            // Profile not found, clear it
            localStorage.removeItem('selected_profile_id')
          }
        }

        // No selected profile, check all profiles
        const profilesResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${user.id}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (profilesResponse.ok) {
          const data = await profilesResponse.json()
          const activeProfiles = data.profiles?.filter((p: any) => p.is_active) || []

          if (activeProfiles.length === 0) {
            // No profiles found, try to sync from existing role records
            const syncResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/sync/${user.id}`, {
              method: 'POST',
              headers: {
                'Authorization': `Bearer ${token}`,
              },
            })

            if (syncResponse.ok) {
              // Profiles were created, get them again
              const profilesResponse2 = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${user.id}`, {
                headers: {
                  'Authorization': `Bearer ${token}`,
                },
              })
              if (profilesResponse2.ok) {
                const profilesData2 = await profilesResponse2.json()
                const syncedProfiles = profilesData2.profiles?.filter((p: any) => p.is_active) || []
                if (syncedProfiles.length > 0) {
                  // Use synced profiles
                  activeProfiles.push(...syncedProfiles)
                }
              }
            }

            if (activeProfiles.length === 0) {
              // Still no profiles, check for existing role records (fallback)
              await checkExistingRolesAndRedirect(token, user.id, businessServiceUrl)
              return
            }
          }

          // Check for default profile first
          const defaultProfile = activeProfiles.find((p: any) => p.is_default && p.is_active)
          
          if (defaultProfile) {
            // Use default profile
            localStorage.setItem('selected_profile_id', defaultProfile.id.toString())
            const profileType = defaultProfile.type
            switch (profileType) {
              case 'gym_owner':
                router.push('/gym-owner')
                return
              case 'trainer':
                router.push('/trainer')
                return
              case 'trainee':
                router.push('/trainee')
                return
            }
          } else if (activeProfiles.length === 1) {
            // Only one profile, automatically select it
            localStorage.setItem('selected_profile_id', activeProfiles[0].id.toString())
            const profileType = activeProfiles[0].type
            switch (profileType) {
              case 'gym_owner':
                router.push('/gym-owner')
                return
              case 'trainer':
                router.push('/trainer')
                return
              case 'trainee':
                router.push('/trainee')
                return
            }
          } else {
            // Multiple profiles but no default, redirect to profile selection
            router.push('/select-profile')
            return
          }
        } else {
          // If profile check fails, check for existing roles
          await checkExistingRolesAndRedirect(token, user.id, businessServiceUrl)
          return
        }
      } catch (error) {
        console.error('Error checking user profiles:', error)
        // Try checking existing roles as fallback
        try {
          const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'
          await checkExistingRolesAndRedirect(token, user.id, businessServiceUrl)
        } catch (err) {
          // If all checks fail, stay on home page
        }
      } finally {
        setCheckingRole(false)
      }
    }

    // Helper function to check existing role records (for backward compatibility)
    const checkExistingRolesAndRedirect = async (token: string, userId: string, businessServiceUrl: string) => {
      try {
        // Check if user is a gym owner
        const gymOwnerResponse = await fetch(`${businessServiceUrl}/api/v1/gym-owners/user/${userId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (gymOwnerResponse.ok) {
          router.push('/gym-owner')
          return
        }

        // TODO: Check for trainer and trainee when those endpoints are available
        // For now, if no gym owner found, redirect to role selection
        router.push('/select-role')
      } catch (error) {
        console.error('Error checking existing roles:', error)
        router.push('/select-role')
      }
    }

    checkUserProfiles()
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

