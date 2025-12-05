'use client'

import { useState, useEffect } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import Link from 'next/link'
import { useAuth } from '@/hooks/useAuth'

interface Profile {
  id: number
  type: 'gym_owner' | 'trainer' | 'trainee'
  person_id: number
  is_active: boolean
  is_default: boolean
}

export default function SwitchProfilePage() {
  const router = useRouter()
  const searchParams = useSearchParams()
  const { user, isAuthenticated, setProfile } = useAuth()
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [checking, setChecking] = useState(false)

  const profileType = searchParams.get('type') as 'trainer' | 'trainee' | null

  useEffect(() => {
    if (!isAuthenticated || !user?.id) {
      router.push('/auth/signin')
      return
    }

    if (!profileType || (profileType !== 'trainer' && profileType !== 'trainee')) {
      setError('Invalid profile type. Please select trainer or trainee.')
      setLoading(false)
      return
    }

    checkAndSwitchProfile()
  }, [isAuthenticated, user, profileType, router])

  const checkAndSwitchProfile = async () => {
    setChecking(true)
    try {
      const token = localStorage.getItem('auth_token')
      if (!token) {
        router.push('/auth/signin')
        return
      }

      const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'

      // Check if profile exists
      const profilesResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${user.id}`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (profilesResponse.ok) {
        const data = await profilesResponse.json()
        const profiles = data.profiles || []
        const existingProfile = profiles.find((p: Profile) => p.type === profileType && p.is_active)

        if (existingProfile) {
          // Profile exists, switch to it
          setProfile(existingProfile.id.toString())
          router.push(`/${profileType}`)
          return
        }
      }

      // Profile doesn't exist, redirect to registration
      router.push(`/register/${profileType}`)
    } catch (err: any) {
      console.error('Error checking profile:', err)
      setError(err.message || 'Failed to check profile')
      setLoading(false)
    } finally {
      setChecking(false)
    }
  }

  if (loading || checking) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <div className="text-center">
          <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-600">Checking profile...</p>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="flex items-center justify-center min-h-[60vh]">
        <div className="max-w-md w-full bg-white rounded-lg shadow-md p-8 text-center">
          <div className="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center mx-auto mb-4">
            <svg className="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
            </svg>
          </div>
          <h2 className="text-xl font-semibold text-gray-900 mb-2">Error</h2>
          <p className="text-gray-600 mb-4">{error}</p>
          <Link
            href="/home"
            className="inline-block px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
          >
            Back to Home
          </Link>
        </div>
      </div>
    )
  }

  return null
}

