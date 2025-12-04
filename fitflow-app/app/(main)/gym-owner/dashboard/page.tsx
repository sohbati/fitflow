'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'

interface GymOwnerData {
  id: number
  person_id: number
  gym_id: number
  brief_bio?: string
  person: {
    id: number
    first_name: string
    last_name: string
    email?: string
    phone_number?: string
    profile_image_url?: string
  }
  gym: {
    id: number
    name: string
    description?: string
    phone_number?: string
    email?: string
    website_url?: string
    is_verified: boolean
  }
}

export default function GymOwnerDashboardPage() {
  const router = useRouter()
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [gymOwnerData, setGymOwnerData] = useState<GymOwnerData | null>(null)

  useEffect(() => {
    const fetchGymOwnerData = async () => {
      try {
        // Get user data from localStorage
        const userDataStr = localStorage.getItem('user_data')
        if (!userDataStr) {
          router.push('/auth/signin')
          return
        }

        const userData = JSON.parse(userDataStr)
        const userId = userData.id
        const token = localStorage.getItem('auth_token')

        if (!token) {
          router.push('/auth/signin')
          return
        }

        // Get business service URL
        const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'

        const response = await fetch(`${businessServiceUrl}/api/v1/gym-owners/user/${userId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (!response.ok) {
          if (response.status === 404) {
            // User is not a gym owner, redirect to home
            router.push('/home')
            return
          }
          throw new Error('Failed to fetch gym owner data')
        }

        const data = await response.json()
        setGymOwnerData(data)
      } catch (err: any) {
        setError(err.message || 'Failed to load gym owner data')
      } finally {
        setLoading(false)
      }
    }

    fetchGymOwnerData()
  }, [router])

  if (loading) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div className="text-center">
          <div className="w-8 h-8 border-4 border-blue-600 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
          <p className="text-gray-600">Loading dashboard...</p>
        </div>
      </div>
    )
  }

  if (error) {
    return (
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
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

  if (!gymOwnerData) {
    return null
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Header */}
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-4">
          <div className="flex justify-between items-center">
            <Link href="/home" className="text-2xl font-bold text-blue-600">
              FitFlow
            </Link>
            <nav className="flex space-x-4">
              <Link
                href="/home"
                className="text-gray-600 hover:text-gray-900 px-3 py-2 rounded-md text-sm font-medium"
              >
                Home
              </Link>
            </nav>
          </div>
        </div>
      </header>

      {/* Main Content */}
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Welcome Section */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex items-center space-x-4">
            {gymOwnerData.person.profile_image_url ? (
              <img
                src={gymOwnerData.person.profile_image_url}
                alt={`${gymOwnerData.person.first_name} ${gymOwnerData.person.last_name}`}
                className="w-16 h-16 rounded-full"
              />
            ) : (
              <div className="w-16 h-16 rounded-full bg-blue-100 flex items-center justify-center">
                <span className="text-2xl font-semibold text-blue-600">
                  {gymOwnerData.person.first_name[0]}{gymOwnerData.person.last_name[0]}
                </span>
              </div>
            )}
            <div>
              <h1 className="text-2xl font-bold text-gray-900">
                Welcome, {gymOwnerData.person.first_name} {gymOwnerData.person.last_name}!
              </h1>
              <p className="text-gray-600">Gym Owner Dashboard</p>
            </div>
          </div>
        </div>

        {/* Gym Information Card */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex items-center justify-between mb-4">
            <h2 className="text-xl font-semibold text-gray-900">Gym Information</h2>
            <span className={`px-3 py-1 rounded-full text-sm font-medium ${
              gymOwnerData.gym.is_verified
                ? 'bg-green-100 text-green-800'
                : 'bg-yellow-100 text-yellow-800'
            }`}>
              {gymOwnerData.gym.is_verified ? 'Verified' : 'Pending Verification'}
            </span>
          </div>
          <div className="space-y-4">
            <div>
              <h3 className="text-lg font-semibold text-gray-900">{gymOwnerData.gym.name}</h3>
              {gymOwnerData.gym.description && (
                <p className="text-gray-600 mt-1">{gymOwnerData.gym.description}</p>
              )}
            </div>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {gymOwnerData.gym.phone_number && (
                <div>
                  <p className="text-sm font-medium text-gray-500">Phone</p>
                  <p className="text-gray-900">{gymOwnerData.gym.phone_number}</p>
                </div>
              )}
              {gymOwnerData.gym.email && (
                <div>
                  <p className="text-sm font-medium text-gray-500">Email</p>
                  <p className="text-gray-900">{gymOwnerData.gym.email}</p>
                </div>
              )}
              {gymOwnerData.gym.website_url && (
                <div>
                  <p className="text-sm font-medium text-gray-500">Website</p>
                  <a
                    href={gymOwnerData.gym.website_url}
                    target="_blank"
                    rel="noopener noreferrer"
                    className="text-blue-600 hover:underline"
                  >
                    {gymOwnerData.gym.website_url}
                  </a>
                </div>
              )}
            </div>
          </div>
        </div>

        {/* Personal Information Card */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Personal Information</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <p className="text-sm font-medium text-gray-500">Full Name</p>
              <p className="text-gray-900">
                {gymOwnerData.person.first_name} {gymOwnerData.person.last_name}
              </p>
            </div>
            {gymOwnerData.person.email && (
              <div>
                <p className="text-sm font-medium text-gray-500">Email</p>
                <p className="text-gray-900">{gymOwnerData.person.email}</p>
              </div>
            )}
            {gymOwnerData.person.phone_number && (
              <div>
                <p className="text-sm font-medium text-gray-500">Phone</p>
                <p className="text-gray-900">{gymOwnerData.person.phone_number}</p>
              </div>
            )}
          </div>
          {gymOwnerData.brief_bio && (
            <div className="mt-4">
              <p className="text-sm font-medium text-gray-500 mb-1">Bio</p>
              <p className="text-gray-900">{gymOwnerData.brief_bio}</p>
            </div>
          )}
        </div>

        {/* Quick Actions */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Quick Actions</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-colors text-left">
              <div className="text-2xl mb-2">👥</div>
              <h3 className="font-semibold text-gray-900">Manage Trainers</h3>
              <p className="text-sm text-gray-600">Add or remove trainers</p>
            </button>
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-colors text-left">
              <div className="text-2xl mb-2">📍</div>
              <h3 className="font-semibold text-gray-900">Manage Locations</h3>
              <p className="text-sm text-gray-600">Add or update gym locations</p>
            </button>
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-blue-500 hover:bg-blue-50 transition-colors text-left">
              <div className="text-2xl mb-2">⚙️</div>
              <h3 className="font-semibold text-gray-900">Gym Settings</h3>
              <p className="text-sm text-gray-600">Update gym information</p>
            </button>
          </div>
        </div>
      </main>
    </div>
  )
}

