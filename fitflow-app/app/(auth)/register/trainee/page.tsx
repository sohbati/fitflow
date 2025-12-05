'use client'

import { useState, FormEvent } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'

interface FormData {
  first_name: string
  last_name: string
  email: string
  phone_number: string
  date_of_birth: string
  gender: string
  address: string
  city: string
  province: string
  country: string
  postal_code: string
  height_cm: string
  weight_kg: string
  fitness_level: string
  goals: string
  medical_conditions: string
  membership_type: string
  membership_start_date: string
  membership_end_date: string
}

export default function TraineeRegistrationPage() {
  const router = useRouter()
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)
  const [formData, setFormData] = useState<FormData>({
    first_name: '',
    last_name: '',
    email: '',
    phone_number: '',
    date_of_birth: '',
    gender: '',
    address: '',
    city: '',
    province: '',
    country: '',
    postal_code: '',
    height_cm: '',
    weight_kg: '',
    fitness_level: 'beginner',
    goals: '',
    medical_conditions: '',
    membership_type: 'basic',
    membership_start_date: '',
    membership_end_date: '',
  })

  const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement | HTMLSelectElement>) => {
    const { name, value } = e.target
    setFormData(prev => ({ ...prev, [name]: value }))
  }

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault()
    setLoading(true)
    setError(null)

    try {
      const token = localStorage.getItem('auth_token')
      if (!token) {
        throw new Error('Not authenticated. Please sign in again.')
      }

      const userDataStr = localStorage.getItem('user_data')
      if (!userDataStr) {
        throw new Error('User data not found. Please sign in again.')
      }

      const userData = JSON.parse(userDataStr)
      const userId = userData.id

      if (!userId || typeof userId !== 'string') {
        throw new Error('Invalid user ID. Please sign in again.')
      }

      const uuidRegex = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i
      if (!uuidRegex.test(userId)) {
        throw new Error('Invalid user ID format. Please sign in again.')
      }

      if (!formData.first_name || !formData.last_name) {
        throw new Error('Please fill in all required fields: First Name and Last Name')
      }

      const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'

      const payload: any = {
        user_id: userId,
        first_name: formData.first_name,
        last_name: formData.last_name,
      }

      if (formData.email && formData.email.trim()) payload.email = formData.email.trim()
      if (formData.phone_number && formData.phone_number.trim()) payload.phone_number = formData.phone_number.trim()
      if (formData.date_of_birth && formData.date_of_birth.trim()) payload.date_of_birth = formData.date_of_birth.trim()
      if (formData.gender && formData.gender.trim()) payload.gender = formData.gender.trim()
      if (formData.address && formData.address.trim()) payload.address = formData.address.trim()
      if (formData.city && formData.city.trim()) payload.city = formData.city.trim()
      if (formData.province && formData.province.trim()) payload.province = formData.province.trim()
      if (formData.country && formData.country.trim()) payload.country = formData.country.trim()
      if (formData.postal_code && formData.postal_code.trim()) payload.postal_code = formData.postal_code.trim()

      // Trainee-specific fields
      if (formData.height_cm && formData.height_cm.trim()) {
        const height = parseInt(formData.height_cm)
        if (!isNaN(height)) payload.height_cm = height
      }
      if (formData.weight_kg && formData.weight_kg.trim()) {
        const weight = parseFloat(formData.weight_kg)
        if (!isNaN(weight)) payload.weight_kg = weight
      }
      if (formData.fitness_level && formData.fitness_level.trim()) payload.fitness_level = formData.fitness_level.trim()
      if (formData.goals && formData.goals.trim()) payload.goals = formData.goals.trim()
      if (formData.medical_conditions && formData.medical_conditions.trim()) payload.medical_conditions = formData.medical_conditions.trim()
      if (formData.membership_type && formData.membership_type.trim()) payload.membership_type = formData.membership_type.trim()
      if (formData.membership_start_date && formData.membership_start_date.trim()) payload.membership_start_date = formData.membership_start_date.trim()
      if (formData.membership_end_date && formData.membership_end_date.trim()) payload.membership_end_date = formData.membership_end_date.trim()

      const response = await fetch(`${businessServiceUrl}/api/v1/persons/register/trainee`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${token}`,
        },
        body: JSON.stringify(payload),
      })

      if (!response.ok) {
        let errorMessage = `Registration failed with status ${response.status}`
        try {
          const errorData = await response.json()
          errorMessage = errorData.error || errorMessage
        } catch (parseError) {
          const errorText = await response.text().catch(() => 'Unknown error')
          errorMessage = errorText || errorMessage
        }
        throw new Error(errorMessage)
      }

      router.push('/trainee/dashboard')
    } catch (err: any) {
      setError(err.message || 'Registration failed. Please try again.')
      setLoading(false)
    }
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-4xl mx-auto">
        <div className="text-center mb-8">
          <Link href="/home" className="text-3xl font-bold text-blue-600">
            FitFlow
          </Link>
          <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
            Register as Trainee
          </h2>
          <p className="mt-2 text-sm text-gray-600">
            Complete your profile to start your fitness journey
          </p>
        </div>

        {error && (
          <div className="mb-6 bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-md p-8 space-y-8">
          {/* Personal Information */}
          <div>
            <h3 className="text-xl font-semibold text-gray-900 mb-4">Personal Information</h3>
            <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
              <div>
                <label htmlFor="first_name" className="block text-sm font-medium text-gray-700">
                  First Name *
                </label>
                <input
                  type="text"
                  name="first_name"
                  id="first_name"
                  required
                  value={formData.first_name}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="last_name" className="block text-sm font-medium text-gray-700">
                  Last Name *
                </label>
                <input
                  type="text"
                  name="last_name"
                  id="last_name"
                  required
                  value={formData.last_name}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="email" className="block text-sm font-medium text-gray-700">
                  Email
                </label>
                <input
                  type="email"
                  name="email"
                  id="email"
                  value={formData.email}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="phone_number" className="block text-sm font-medium text-gray-700">
                  Phone Number
                </label>
                <input
                  type="tel"
                  name="phone_number"
                  id="phone_number"
                  value={formData.phone_number}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="date_of_birth" className="block text-sm font-medium text-gray-700">
                  Date of Birth
                </label>
                <input
                  type="date"
                  name="date_of_birth"
                  id="date_of_birth"
                  value={formData.date_of_birth}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="gender" className="block text-sm font-medium text-gray-700">
                  Gender
                </label>
                <select
                  name="gender"
                  id="gender"
                  value={formData.gender}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                  <option value="">Select...</option>
                  <option value="male">Male</option>
                  <option value="female">Female</option>
                  <option value="other">Other</option>
                </select>
              </div>
            </div>
          </div>

          {/* Fitness Information */}
          <div>
            <h3 className="text-xl font-semibold text-gray-900 mb-4">Fitness Information</h3>
            <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
              <div>
                <label htmlFor="height_cm" className="block text-sm font-medium text-gray-700">
                  Height (cm)
                </label>
                <input
                  type="number"
                  name="height_cm"
                  id="height_cm"
                  value={formData.height_cm}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="weight_kg" className="block text-sm font-medium text-gray-700">
                  Weight (kg)
                </label>
                <input
                  type="number"
                  step="0.1"
                  name="weight_kg"
                  id="weight_kg"
                  value={formData.weight_kg}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="fitness_level" className="block text-sm font-medium text-gray-700">
                  Fitness Level
                </label>
                <select
                  name="fitness_level"
                  id="fitness_level"
                  value={formData.fitness_level}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                  <option value="beginner">Beginner</option>
                  <option value="intermediate">Intermediate</option>
                  <option value="advanced">Advanced</option>
                </select>
              </div>
              <div>
                <label htmlFor="membership_type" className="block text-sm font-medium text-gray-700">
                  Membership Type
                </label>
                <select
                  name="membership_type"
                  id="membership_type"
                  value={formData.membership_type}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                >
                  <option value="basic">Basic</option>
                  <option value="premium">Premium</option>
                  <option value="vip">VIP</option>
                </select>
              </div>
              <div className="sm:col-span-2">
                <label htmlFor="goals" className="block text-sm font-medium text-gray-700">
                  Fitness Goals
                </label>
                <textarea
                  name="goals"
                  id="goals"
                  rows={3}
                  value={formData.goals}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div className="sm:col-span-2">
                <label htmlFor="medical_conditions" className="block text-sm font-medium text-gray-700">
                  Medical Conditions
                </label>
                <textarea
                  name="medical_conditions"
                  id="medical_conditions"
                  rows={3}
                  value={formData.medical_conditions}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                  placeholder="Please list any medical conditions or injuries..."
                />
              </div>
              <div>
                <label htmlFor="membership_start_date" className="block text-sm font-medium text-gray-700">
                  Membership Start Date
                </label>
                <input
                  type="date"
                  name="membership_start_date"
                  id="membership_start_date"
                  value={formData.membership_start_date}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="membership_end_date" className="block text-sm font-medium text-gray-700">
                  Membership End Date
                </label>
                <input
                  type="date"
                  name="membership_end_date"
                  id="membership_end_date"
                  value={formData.membership_end_date}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
            </div>
          </div>

          <div className="flex justify-end space-x-4">
            <Link
              href="/home"
              className="px-6 py-2 border border-gray-300 rounded-md text-gray-700 hover:bg-gray-50"
            >
              Cancel
            </Link>
            <button
              type="submit"
              disabled={loading}
              className="px-6 py-2 bg-purple-600 text-white rounded-md hover:bg-purple-700 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {loading ? 'Registering...' : 'Register as Trainee'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

