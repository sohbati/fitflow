'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'

interface TrainerData {
  id: number
  person_id: number
  is_registered: boolean
  person: {
    id: number
    first_name: string
    last_name: string
    email?: string
    phone_number?: string
    profile_image_url?: string
  }
  gyms?: Array<{
    id: number
    name: string
    description?: string
  }>
}

export default function TrainerDashboardPage() {
  const router = useRouter()
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [trainerData, setTrainerData] = useState<TrainerData | null>(null)

  useEffect(() => {
    const fetchTrainerData = async () => {
      try {
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

        const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'

        const response = await fetch(`${businessServiceUrl}/api/v1/trainers/user/${userId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (!response.ok) {
          if (response.status === 404) {
            router.push('/home')
            return
          }
          throw new Error('Failed to fetch trainer data')
        }

        const data = await response.json()
        setTrainerData(data)
      } catch (err: any) {
        setError(err.message || 'Failed to load trainer data')
      } finally {
        setLoading(false)
      }
    }

    fetchTrainerData()
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

  if (!trainerData) {
    return null
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Welcome Section */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex items-center space-x-4">
            {trainerData.person.profile_image_url ? (
              <img
                src={trainerData.person.profile_image_url}
                alt={`${trainerData.person.first_name} ${trainerData.person.last_name}`}
                className="w-16 h-16 rounded-full"
              />
            ) : (
              <div className="w-16 h-16 rounded-full bg-green-100 flex items-center justify-center">
                <span className="text-2xl font-semibold text-green-600">
                  {trainerData.person.first_name[0]}{trainerData.person.last_name[0]}
                </span>
              </div>
            )}
            <div>
              <h1 className="text-2xl font-bold text-gray-900">
                Welcome, {trainerData.person.first_name} {trainerData.person.last_name}!
              </h1>
              <p className="text-gray-600">Trainer Dashboard</p>
            </div>
          </div>
        </div>

        {/* Status Card */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex items-center justify-between">
            <h2 className="text-xl font-semibold text-gray-900">Registration Status</h2>
            <span className={`px-3 py-1 rounded-full text-sm font-medium ${
              trainerData.is_registered
                ? 'bg-green-100 text-green-800'
                : 'bg-yellow-100 text-yellow-800'
            }`}>
              {trainerData.is_registered ? 'Registered' : 'Not Registered'}
            </span>
          </div>
          {!trainerData.is_registered && (
            <p className="text-gray-600 mt-2">Register with a gym to start training clients.</p>
          )}
        </div>

        {/* Gyms Section */}
        {trainerData.gyms && trainerData.gyms.length > 0 && (
          <div className="bg-white rounded-lg shadow-md p-6 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Associated Gyms</h2>
            <div className="space-y-4">
              {trainerData.gyms.map((gym) => (
                <div key={gym.id} className="border border-gray-200 rounded-lg p-4">
                  <h3 className="font-semibold text-gray-900">{gym.name}</h3>
                  {gym.description && (
                    <p className="text-gray-600 text-sm mt-1">{gym.description}</p>
                  )}
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Personal Information Card */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Personal Information</h2>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <p className="text-sm font-medium text-gray-500">Full Name</p>
              <p className="text-gray-900">
                {trainerData.person.first_name} {trainerData.person.last_name}
              </p>
            </div>
            {trainerData.person.email && (
              <div>
                <p className="text-sm font-medium text-gray-500">Email</p>
                <p className="text-gray-900">{trainerData.person.email}</p>
              </div>
            )}
            {trainerData.person.phone_number && (
              <div>
                <p className="text-sm font-medium text-gray-500">Phone</p>
                <p className="text-gray-900">{trainerData.person.phone_number}</p>
              </div>
            )}
          </div>
        </div>

        {/* Quick Actions */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Quick Actions</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-green-500 hover:bg-green-50 transition-colors text-left">
              <div className="text-2xl mb-2">👥</div>
              <h3 className="font-semibold text-gray-900">My Clients</h3>
              <p className="text-sm text-gray-600">View and manage clients</p>
            </button>
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-green-500 hover:bg-green-50 transition-colors text-left">
              <div className="text-2xl mb-2">📅</div>
              <h3 className="font-semibold text-gray-900">Schedule</h3>
              <p className="text-sm text-gray-600">Manage training sessions</p>
            </button>
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-green-500 hover:bg-green-50 transition-colors text-left">
              <div className="text-2xl mb-2">📊</div>
              <h3 className="font-semibold text-gray-900">Progress</h3>
              <p className="text-sm text-gray-600">Track client progress</p>
            </button>
          </div>
        </div>
      </main>
    </div>
  )
}

