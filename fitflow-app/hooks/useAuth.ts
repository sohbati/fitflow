'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'

export interface User {
  id?: string
  email?: string
  name?: string
  picture?: string
  google_id?: string
  mobile?: string
  country?: string
}

export function useAuth() {
  const router = useRouter()
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState<User | null>(null)
  const [selectedProfileId, setSelectedProfileId] = useState<string | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    checkAuth()
  }, [])

  const checkAuth = () => {
    const token = localStorage.getItem('auth_token')
    const userData = localStorage.getItem('user_data')
    const profileId = localStorage.getItem('selected_profile_id')
    
    if (token && userData) {
      try {
        setIsAuthenticated(true)
        setUser(JSON.parse(userData))
        setSelectedProfileId(profileId)
      } catch (error) {
        console.error('Error parsing user data:', error)
        clearAuth()
      }
    } else {
      setIsAuthenticated(false)
      setUser(null)
      setSelectedProfileId(null)
    }
    setIsLoading(false)
  }

  const clearAuth = () => {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
    localStorage.removeItem('selected_profile_id')
    setIsAuthenticated(false)
    setUser(null)
    setSelectedProfileId(null)
  }

  const signOut = () => {
    clearAuth()
    // Redirect to sign-in page after signing out
    router.push('/auth/signin')
  }

  const setProfile = (profileId: string) => {
    localStorage.setItem('selected_profile_id', profileId)
    setSelectedProfileId(profileId)
  }

  return {
    isAuthenticated,
    user,
    selectedProfileId,
    isLoading,
    signOut,
    checkAuth,
    setProfile,
  }
}

