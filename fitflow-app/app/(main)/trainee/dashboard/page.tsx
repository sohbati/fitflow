'use client'

import { useState, useEffect } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'

interface TraineeData {
  id: number
  person_id: number
  height_cm?: number
  weight_kg?: number
  fitness_level: 'beginner' | 'intermediate' | 'advanced'
  goals?: string
  medical_conditions?: string
  membership_type: 'basic' | 'premium' | 'vip'
  membership_start_date?: string
  membership_end_date?: string
  is_active: boolean
  person: {
    id: number
    first_name: string
    last_name: string
    email?: string
    phone_number?: string
    profile_image_url?: string
    date_of_birth?: string
  }
}

export default function TraineeDashboardPage() {
  const router = useRouter()
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)
  const [traineeData, setTraineeData] = useState<TraineeData | null>(null)

  useEffect(() => {
    const fetchTraineeData = async () => {
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

        const response = await fetch(`${businessServiceUrl}/api/v1/trainees/user/${userId}`, {
          headers: {
            'Authorization': `Bearer ${token}`,
          },
        })

        if (!response.ok) {
          if (response.status === 404) {
            router.push('/home')
            return
          }
          throw new Error('Failed to fetch trainee data')
        }

        const data = await response.json()
        setTraineeData(data)
      } catch (err: any) {
        setError(err.message || 'Failed to load trainee data')
      } finally {
        setLoading(false)
      }
    }

    fetchTraineeData()
  }, [router])

  const calculateBMI = () => {
    if (!traineeData?.height_cm || !traineeData?.weight_kg) return null
    const heightM = traineeData.height_cm / 100
    return (traineeData.weight_kg / (heightM * heightM)).toFixed(1)
  }

  const getFitnessLevelColor = (level: string) => {
    switch (level) {
      case 'beginner':
        return 'bg-blue-100 text-blue-800'
      case 'intermediate':
        return 'bg-yellow-100 text-yellow-800'
      case 'advanced':
        return 'bg-green-100 text-green-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

  const getMembershipColor = (type: string) => {
    switch (type) {
      case 'basic':
        return 'bg-gray-100 text-gray-800'
      case 'premium':
        return 'bg-blue-100 text-blue-800'
      case 'vip':
        return 'bg-purple-100 text-purple-800'
      default:
        return 'bg-gray-100 text-gray-800'
    }
  }

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

  if (!traineeData) {
    return null
  }

  const bmi = calculateBMI()

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Welcome Section */}
        <div className="bg-white rounded-lg shadow-md p-6 mb-6">
          <div className="flex items-center space-x-4">
            {traineeData.person.profile_image_url ? (
              <img
                src={traineeData.person.profile_image_url}
                alt={`${traineeData.person.first_name} ${traineeData.person.last_name}`}
                className="w-16 h-16 rounded-full"
              />
            ) : (
              <div className="w-16 h-16 rounded-full bg-purple-100 flex items-center justify-center">
                <span className="text-2xl font-semibold text-purple-600">
                  {traineeData.person.first_name[0]}{traineeData.person.last_name[0]}
                </span>
              </div>
            )}
            <div>
              <h1 className="text-2xl font-bold text-gray-900">
                Welcome, {traineeData.person.first_name} {traineeData.person.last_name}!
              </h1>
              <p className="text-gray-600">Trainee Dashboard</p>
            </div>
          </div>
        </div>

        {/* Stats Cards */}
        <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-6">
          {/* Fitness Level */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-500">Fitness Level</p>
                <p className="text-2xl font-bold text-gray-900 mt-1 capitalize">
                  {traineeData.fitness_level}
                </p>
              </div>
              <span className={`px-3 py-1 rounded-full text-sm font-medium ${getFitnessLevelColor(traineeData.fitness_level)}`}>
                {traineeData.fitness_level}
              </span>
            </div>
          </div>

          {/* Membership */}
          <div className="bg-white rounded-lg shadow-md p-6">
            <div className="flex items-center justify-between">
              <div>
                <p className="text-sm font-medium text-gray-500">Membership</p>
                <p className="text-2xl font-bold text-gray-900 mt-1 capitalize">
                  {traineeData.membership_type}
                </p>
              </div>
              <span className={`px-3 py-1 rounded-full text-sm font-medium ${getMembershipColor(traineeData.membership_type)}`}>
                {traineeData.membership_type}
              </span>
            </div>
          </div>

          {/* BMI */}
          {bmi && (
            <div className="bg-white rounded-lg shadow-md p-6">
              <div>
                <p className="text-sm font-medium text-gray-500">BMI</p>
                <p className="text-2xl font-bold text-gray-900 mt-1">{bmi}</p>
              </div>
            </div>
          )}
        </div>

        {/* Body Metrics */}
        {(traineeData.height_cm || traineeData.weight_kg) && (
          <div className="bg-white rounded-lg shadow-md p-6 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Body Metrics</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {traineeData.height_cm && (
                <div>
                  <p className="text-sm font-medium text-gray-500">Height</p>
                  <p className="text-gray-900">{traineeData.height_cm} cm</p>
                </div>
              )}
              {traineeData.weight_kg && (
                <div>
                  <p className="text-sm font-medium text-gray-500">Weight</p>
                  <p className="text-gray-900">{traineeData.weight_kg} kg</p>
                </div>
              )}
            </div>
          </div>
        )}

        {/* Goals */}
        {traineeData.goals && (
          <div className="bg-white rounded-lg shadow-md p-6 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Fitness Goals</h2>
            <p className="text-gray-700">{traineeData.goals}</p>
          </div>
        )}

        {/* Medical Conditions */}
        {traineeData.medical_conditions && (
          <div className="bg-white rounded-lg shadow-md p-6 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Medical Conditions</h2>
            <p className="text-gray-700">{traineeData.medical_conditions}</p>
          </div>
        )}

        {/* Membership Info */}
        {(traineeData.membership_start_date || traineeData.membership_end_date) && (
          <div className="bg-white rounded-lg shadow-md p-6 mb-6">
            <h2 className="text-xl font-semibold text-gray-900 mb-4">Membership Information</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
              {traineeData.membership_start_date && (
                <div>
                  <p className="text-sm font-medium text-gray-500">Start Date</p>
                  <p className="text-gray-900">{new Date(traineeData.membership_start_date).toLocaleDateString()}</p>
                </div>
              )}
              {traineeData.membership_end_date && (
                <div>
                  <p className="text-sm font-medium text-gray-500">End Date</p>
                  <p className="text-gray-900">{new Date(traineeData.membership_end_date).toLocaleDateString()}</p>
                </div>
              )}
            </div>
            <div className="mt-4">
              <span className={`px-3 py-1 rounded-full text-sm font-medium ${
                traineeData.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'
              }`}>
                {traineeData.is_active ? 'Active' : 'Inactive'}
              </span>
            </div>
          </div>
        )}

        {/* Quick Actions */}
        <div className="bg-white rounded-lg shadow-md p-6">
          <h2 className="text-xl font-semibold text-gray-900 mb-4">Quick Actions</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-purple-500 hover:bg-purple-50 transition-colors text-left">
              <div className="text-2xl mb-2">💪</div>
              <h3 className="font-semibold text-gray-900">My Workouts</h3>
              <p className="text-sm text-gray-600">View workout history</p>
            </button>
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-purple-500 hover:bg-purple-50 transition-colors text-left">
              <div className="text-2xl mb-2">📊</div>
              <h3 className="font-semibold text-gray-900">Progress</h3>
              <p className="text-sm text-gray-600">Track your fitness journey</p>
            </button>
            <button className="p-4 border-2 border-gray-200 rounded-lg hover:border-purple-500 hover:bg-purple-50 transition-colors text-left">
              <div className="text-2xl mb-2">👨‍🏫</div>
              <h3 className="font-semibold text-gray-900">My Trainer</h3>
              <p className="text-sm text-gray-600">Connect with your trainer</p>
            </button>
          </div>
        </div>
      </main>
    </div>
  )
}

