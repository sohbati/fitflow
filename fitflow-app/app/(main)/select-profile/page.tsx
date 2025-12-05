'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { useAuth } from '@/hooks/useAuth'

interface Profile {
  id: number
  type: 'gym_owner' | 'trainer' | 'trainee'
  person_id: number
  is_active: boolean
  is_default: boolean
  person?: {
    id: number
    first_name: string
    last_name: string
    email?: string
  }
}

export default function SelectProfilePage() {
  const router = useRouter()
  const { user, isAuthenticated, setProfile, isLoading } = useAuth()
  const [profiles, setProfiles] = useState<Profile[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [selecting, setSelecting] = useState(false)

  useEffect(() => {
    // Wait for auth hook to finish loading
    if (isLoading) {
      return
    }

    // Check authentication directly from localStorage as well
    const token = localStorage.getItem('auth_token')
    const userData = localStorage.getItem('user_data')
    
    if (!token || !userData) {
      router.push('/auth/signin')
      return
    }

    // Parse user data to get user ID
    let userId: string | null = null
    try {
      const parsedUser = JSON.parse(userData)
      userId = parsedUser?.id
    } catch (e) {
      console.error('Error parsing user data:', e)
      router.push('/auth/signin')
      return
    }

    if (!userId) {
      router.push('/auth/signin')
      return
    }

    // If we have token and user ID, fetch profiles
    fetchProfiles(userId)
  }, [isLoading, isAuthenticated, user, router])

  const fetchProfiles = async (userId: string) => {
    try {
      const token = localStorage.getItem('auth_token')
      if (!token) {
        router.push('/auth/signin')
        return
      }

      const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'
      
      const response = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${userId}`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (!response.ok) {
        throw new Error('Failed to fetch profiles')
      }

      const data = await response.json()
      const activeProfiles = data.profiles?.filter((p: Profile) => p.is_active) || []

      if (activeProfiles.length === 0) {
        // No profiles, redirect to role selection
        router.push('/select-role')
        return
      }

      if (activeProfiles.length === 1) {
        // Only one profile, automatically select it
        await selectProfile(activeProfiles[0].id)
        return
      }

      setProfiles(activeProfiles)
    } catch (err: any) {
      console.error('Error fetching profiles:', err)
      setError(err.message || 'Failed to load profiles')
    } finally {
      setLoading(false)
    }
  }

  const selectProfile = async (profileId: number) => {
    setSelecting(true)
    try {
      // Store selected profile using useAuth hook
      setProfile(profileId.toString())
      
      // Get the selected profile to determine redirect
      const selectedProfile = profiles.find(p => p.id === profileId)
      
      // Redirect based on profile type
      switch (selectedProfile?.type) {
        case 'gym_owner':
          router.push('/gym-owner')
          break
        case 'trainer':
          router.push('/trainer')
          break
        case 'trainee':
          router.push('/trainee')
          break
        default:
          router.push('/home')
      }
    } catch (err: any) {
      console.error('Error selecting profile:', err)
      setError(err.message || 'Failed to select profile')
    } finally {
      setSelecting(false)
    }
  }

  const getProfileTypeLabel = (type: string) => {
    switch (type) {
      case 'gym_owner':
        return 'Gym Owner'
      case 'trainer':
        return 'Trainer'
      case 'trainee':
        return 'Trainee'
      default:
        return type
    }
  }

  const getProfileTypeDescription = (type: string) => {
    switch (type) {
      case 'gym_owner':
        return 'Manage your gym, locations, and trainers'
      case 'trainer':
        return 'Train clients and manage your schedule'
      case 'trainee':
        return 'Track your workouts and progress'
      default:
        return ''
    }
  }

  const getProfileIcon = (type: string) => {
    switch (type) {
      case 'gym_owner':
        return '🏋️'
      case 'trainer':
        return '👨‍🏫'
      case 'trainee':
        return '💪'
      default:
        return '👤'
    }
  }

  if (isLoading || loading) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <div className="text-center">
          <div className="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto"></div>
          <p className="mt-4 text-gray-600">{isLoading ? 'Checking authentication...' : 'Loading profiles...'}</p>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <div className="text-center">
          <p className="text-red-600 mb-4">{error}</p>
          <button
            onClick={() => router.push('/auth/signin')}
            className="px-4 py-2 bg-blue-600 text-white rounded hover:bg-blue-700"
          >
            Back to Sign In
          </button>
        </div>
      </div>
    )
  }

  return (
    <div className="flex items-center justify-center min-h-[60vh] py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-2xl w-full space-y-8">
        <div className="text-center">
          <Link href="/home" className="text-3xl font-bold text-blue-600">
            FitFlow
          </Link>
          <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
            Select Your Profile
          </h2>
          <p className="mt-2 text-sm text-gray-600">
            You have multiple profiles. Choose which role you want to use:
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-3 gap-4">
          {profiles.map((profile) => (
            <button
              key={profile.id}
              onClick={() => selectProfile(profile.id)}
              disabled={selecting}
              className={`
                relative p-6 rounded-lg border-2 transition-all
                ${profile.is_default 
                  ? 'border-blue-500 bg-blue-50' 
                  : 'border-gray-200 bg-white hover:border-blue-300 hover:bg-blue-50'
                }
                ${selecting ? 'opacity-50 cursor-not-allowed' : 'cursor-pointer'}
              `}
            >
              {profile.is_default && (
                <span className="absolute top-2 right-2 text-xs bg-blue-500 text-white px-2 py-1 rounded">
                  Default
                </span>
              )}
              
              <div className="text-center">
                <div className="text-4xl mb-3">{getProfileIcon(profile.type)}</div>
                <h3 className="text-lg font-semibold text-gray-900 mb-1">
                  {getProfileTypeLabel(profile.type)}
                </h3>
                <p className="text-sm text-gray-600 mb-3">
                  {getProfileTypeDescription(profile.type)}
                </p>
                {profile.person && (
                  <p className="text-xs text-gray-500">
                    {profile.person.first_name} {profile.person.last_name}
                  </p>
                )}
              </div>
            </button>
          ))}
        </div>

        {(() => {
          const hasTrainer = profiles.some(p => p.type === 'trainer')
          const hasTrainee = profiles.some(p => p.type === 'trainee')
          
          // Only show registration buttons for profiles the user doesn't have
          const missingProfiles = []
          if (!hasTrainer) missingProfiles.push({ type: 'trainer', label: 'Register as Trainer', className: 'bg-green-600 hover:bg-green-700' })
          if (!hasTrainee) missingProfiles.push({ type: 'trainee', label: 'Register as Trainee', className: 'bg-purple-600 hover:bg-purple-700' })
          
          // Don't show gym owner registration here since it's a different flow
          
          if (missingProfiles.length > 0) {
            return (
              <div className="text-center space-y-4">
                <p className="text-sm text-gray-600">Don't have a profile yet?</p>
                <div className="flex justify-center space-x-4">
                  {missingProfiles.map((profile) => (
                    <Link
                      key={profile.type}
                      href={`/switch-profile?type=${profile.type}`}
                      className={`px-4 py-2 ${profile.className} text-white rounded-md text-sm`}
                    >
                      {profile.label}
                    </Link>
                  ))}
                </div>
              </div>
            )
          }
          return null
        })()}
      </div>
    </div>
  )
}

