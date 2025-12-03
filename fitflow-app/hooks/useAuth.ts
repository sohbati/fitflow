'use client'

import { useState, useEffect } from 'react'

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
  const [isAuthenticated, setIsAuthenticated] = useState(false)
  const [user, setUser] = useState<User | null>(null)
  const [isLoading, setIsLoading] = useState(true)

  useEffect(() => {
    checkAuth()
  }, [])

  const checkAuth = () => {
    const token = localStorage.getItem('auth_token')
    const userData = localStorage.getItem('user_data')
    
    if (token && userData) {
      try {
        setIsAuthenticated(true)
        setUser(JSON.parse(userData))
      } catch (error) {
        console.error('Error parsing user data:', error)
        clearAuth()
      }
    } else {
      setIsAuthenticated(false)
      setUser(null)
    }
    setIsLoading(false)
  }

  const clearAuth = () => {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
    setIsAuthenticated(false)
    setUser(null)
  }

  const signOut = () => {
    clearAuth()
  }

  return {
    isAuthenticated,
    user,
    isLoading,
    signOut,
    checkAuth,
  }
}

