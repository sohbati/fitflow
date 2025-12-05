'use client'

import { useEffect, useState, Suspense } from 'react'
import { useRouter, useSearchParams } from 'next/navigation'
import Link from 'next/link'
import { handleGoogleCallback } from '@/hooks/useGoogleAuth'

function GoogleCallbackContent() {
  const router = useRouter()
  const searchParams = useSearchParams()
  const [status, setStatus] = useState<'loading' | 'success' | 'error'>('loading')
  const [message, setMessage] = useState('')

  // Helper function to check profiles and redirect
  const checkProfilesAndRedirect = async (token: string, userId: string) => {
    try {
      const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'
      
      // First, check for profiles
      const response = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${userId}`, {
        headers: {
          'Authorization': `Bearer ${token}`,
        },
      })

      if (response.ok) {
        const data = await response.json()
        let activeProfiles = data.profiles?.filter((p: any) => p.is_active) || []

        if (activeProfiles.length === 0) {
          // No profiles found, try to sync from existing role records
          const syncResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/sync/${userId}`, {
            method: 'POST',
            headers: {
              'Authorization': `Bearer ${token}`,
            },
          })

          if (syncResponse.ok) {
            // Profiles were created, get them again
            const profilesResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${userId}`, {
              headers: {
                'Authorization': `Bearer ${token}`,
              },
            })
            if (profilesResponse.ok) {
              const profilesData = await profilesResponse.json()
              activeProfiles = profilesData.profiles?.filter((p: any) => p.is_active) || []
            }
          }

          if (activeProfiles.length === 0) {
            // Still no profiles, redirect to role selection
            router.push('/select-role')
            return
          }
        }

        // Clear any previously selected profile on new login
        localStorage.removeItem('selected_profile_id')

        if (activeProfiles.length === 1) {
          // Only one profile, automatically select it and redirect
          localStorage.setItem('selected_profile_id', activeProfiles[0].id.toString())
          const profileType = activeProfiles[0].type
          switch (profileType) {
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
        } else if (activeProfiles.length > 1) {
          // Multiple profiles, redirect to profile selection
          router.push('/select-profile')
        } else {
          // No active profiles, redirect to role selection
          router.push('/select-role')
        }
      } else {
        // If profile check fails, try to sync first
        console.warn('Profile check failed, attempting to sync profiles...')
        const syncResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/sync/${userId}`, {
          method: 'POST',
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (syncResponse.ok) {
          // Retry profile check after sync
          const retryResponse = await fetch(`${businessServiceUrl}/api/v1/profiles/user/${userId}`, {
            headers: {
              'Authorization': `Bearer ${token}`,
            },
          })
          if (retryResponse.ok) {
            const retryData = await retryResponse.json()
            const retryProfiles = retryData.profiles?.filter((p: any) => p.is_active) || []
            if (retryProfiles.length > 1) {
              router.push('/select-profile')
              return
            } else if (retryProfiles.length === 1) {
              localStorage.setItem('selected_profile_id', retryProfiles[0].id.toString())
              const profileType = retryProfiles[0].type
              switch (profileType) {
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
              return
            }
          }
        }
        
        // If all else fails, check for existing roles or redirect to role selection
        await checkExistingRolesAndRedirect(token, userId, businessServiceUrl)
      }
    } catch (error) {
      console.error('Error checking profiles:', error)
      // On error, redirect to home page which will handle profile checking
      router.push('/home')
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

  useEffect(() => {
    const code = searchParams.get('code')
    const error = searchParams.get('error')

    if (error) {
      setStatus('error')
      setMessage(`Google authentication failed: ${error}`)
      return
    }

    if (code) {
      const authenticateWithGoogle = async () => {
        try {
          const result = await handleGoogleCallback(code)
          
          setStatus('success')
          setMessage('Successfully signed in with Google!')
          
          // Check profiles and redirect accordingly
          setTimeout(() => {
            checkProfilesAndRedirect(result.token, result.user.id)
          }, 1000)
          
        } catch (error: any) {
          console.error('Google authentication error:', error)
          setStatus('error')
          setMessage(`Authentication failed: ${error.message}`)
        }
      }
      
      authenticateWithGoogle()
    } else {
      setStatus('error')
      setMessage('Google authentication failed: No code received.')
    }
  }, [searchParams, router])

  return (
    <div className="flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div className="text-center">
          <Link href="/home" className="text-3xl font-bold text-blue-600">
            FitFlow
          </Link>
        </div>
        
        <div className="bg-white rounded-lg shadow-md p-8 text-center">
          {status === 'loading' && (
            <div className="space-y-4">
              <div className="flex justify-center">
                <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin"></div>
              </div>
              <h2 className="text-xl font-semibold text-gray-900">Authenticating with Google...</h2>
              <p className="text-gray-600">Please wait while we complete your sign-in.</p>
            </div>
          )}
          
          {status === 'success' && (
            <div className="space-y-4">
              <div className="flex justify-center">
                <div className="w-16 h-16 bg-green-100 rounded-full flex items-center justify-center">
                  <svg className="w-8 h-8 text-green-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M5 13l4 4L19 7" />
                  </svg>
                </div>
              </div>
              <h2 className="text-xl font-semibold text-gray-900">Success!</h2>
              <p className="text-gray-600">{message}</p>
              <p className="text-sm text-gray-500">Redirecting to home page...</p>
            </div>
          )}
          
          {status === 'error' && (
            <div className="space-y-4">
              <div className="flex justify-center">
                <div className="w-16 h-16 bg-red-100 rounded-full flex items-center justify-center">
                  <svg className="w-8 h-8 text-red-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M6 18L18 6M6 6l12 12" />
                  </svg>
                </div>
              </div>
              <h2 className="text-xl font-semibold text-gray-900">Authentication Failed</h2>
              <p className="text-gray-600">{message}</p>
              <div className="space-y-2">
                <Link 
                  href="/auth/signin"
                  className="inline-block w-full px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition-colors"
                >
                  Try Again
                </Link>
                <Link 
                  href="/home"
                  className="inline-block w-full px-4 py-2 bg-gray-200 text-gray-800 rounded-md hover:bg-gray-300 transition-colors"
                >
                  Back to Home
                </Link>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  )
}

export default function GoogleCallbackPage() {
  return (
    <Suspense fallback={
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div className="text-center">
          <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-600">Loading...</p>
        </div>
      </div>
    }>
      <GoogleCallbackContent />
    </Suspense>
  )
}

