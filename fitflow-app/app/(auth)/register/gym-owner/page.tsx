'use client'

import { useState, FormEvent } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'

interface FormData {
  // Person fields
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

  // Gym fields
  gym_name: string
  gym_description: string
  gym_phone_number: string
  gym_email: string
  gym_website_url: string

  // Gym owner fields
  brief_bio: string
}

export default function GymOwnerRegistrationPage() {
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
    gym_name: '',
    gym_description: '',
    gym_phone_number: '',
    gym_email: '',
    gym_website_url: '',
    brief_bio: '',
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
      // Get auth token from localStorage
      const token = localStorage.getItem('auth_token')
      if (!token) {
        throw new Error('Not authenticated. Please sign in again.')
      }

      // Get user data to extract user ID
      const userDataStr = localStorage.getItem('user_data')
      if (!userDataStr) {
        throw new Error('User data not found. Please sign in again.')
      }

      const userData = JSON.parse(userDataStr)
      const userId = userData.id

      // Validate user ID format (should be UUID)
      if (!userId || typeof userId !== 'string') {
        throw new Error('Invalid user ID. Please sign in again.')
      }

      // Validate UUID format (basic check)
      const uuidRegex = /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i
      if (!uuidRegex.test(userId)) {
        console.error('Invalid UUID format for user_id:', userId)
        throw new Error('Invalid user ID format. Please sign in again.')
      }

      console.log('User ID:', userId)
      console.log('Form data:', formData)

      // Get business service URL (assuming it's proxied through gateway)
      const businessServiceUrl = process.env.NEXT_PUBLIC_BUSINESS_SERVICE_URL || 'http://localhost:8090'

      // Validate required fields
      if (!formData.first_name || !formData.last_name || !formData.gym_name) {
        throw new Error('Please fill in all required fields: First Name, Last Name, and Gym Name')
      }

      if (!userId) {
        throw new Error('User ID not found. Please sign in again.')
      }

      // Prepare request payload - match backend struct field names exactly
      const payload: any = {
        user_id: userId, // Required: User ID from IAM service
        first_name: formData.first_name, // Required
        last_name: formData.last_name, // Required
        gym_name: formData.gym_name, // Required
      }

      // Add optional fields only if they have values (use null for empty strings to match backend expectations)
      if (formData.email && formData.email.trim()) payload.email = formData.email.trim()
      if (formData.phone_number && formData.phone_number.trim()) payload.phone_number = formData.phone_number.trim()
      if (formData.date_of_birth && formData.date_of_birth.trim()) payload.date_of_birth = formData.date_of_birth.trim()
      if (formData.gender && formData.gender.trim()) payload.gender = formData.gender.trim()
      if (formData.address && formData.address.trim()) payload.address = formData.address.trim()
      if (formData.city && formData.city.trim()) payload.city = formData.city.trim()
      if (formData.province && formData.province.trim()) payload.province = formData.province.trim()
      if (formData.country && formData.country.trim()) payload.country = formData.country.trim()
      if (formData.postal_code && formData.postal_code.trim()) payload.postal_code = formData.postal_code.trim()
      if (formData.gym_description && formData.gym_description.trim()) payload.gym_description = formData.gym_description.trim()
      if (formData.gym_phone_number && formData.gym_phone_number.trim()) payload.gym_phone_number = formData.gym_phone_number.trim()
      if (formData.gym_email && formData.gym_email.trim()) payload.gym_email = formData.gym_email.trim()
      if (formData.gym_website_url && formData.gym_website_url.trim()) payload.gym_website_url = formData.gym_website_url.trim()
      if (formData.brief_bio && formData.brief_bio.trim()) payload.brief_bio = formData.brief_bio.trim()

      console.log('Sending payload:', JSON.stringify(payload, null, 2))

      const response = await fetch(`${businessServiceUrl}/api/v1/persons/register/gym-owner`, {
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
          console.error('Registration error response:', errorData)
          console.error('Request payload:', JSON.stringify(payload, null, 2))
          errorMessage = errorData.error || errorMessage
        } catch (parseError) {
          const errorText = await response.text().catch(() => 'Unknown error')
          console.error('Failed to parse error response:', errorText)
          errorMessage = errorText || errorMessage
        }
        throw new Error(errorMessage)
      }

      // Registration successful, redirect to gym owner dashboard
      router.push('/gym-owner/dashboard')
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
            Register as Gym Owner
          </h2>
          <p className="mt-2 text-sm text-gray-600">
            Complete your profile and gym information
          </p>
        </div>

        {error && (
          <div className="mb-6 bg-red-50 border border-red-200 text-red-700 px-4 py-3 rounded">
            {error}
          </div>
        )}

        <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow-md p-8 space-y-8">
          {/* Personal Information Section */}
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
              <div className="sm:col-span-2">
                <label htmlFor="address" className="block text-sm font-medium text-gray-700">
                  Address
                </label>
                <input
                  type="text"
                  name="address"
                  id="address"
                  value={formData.address}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="city" className="block text-sm font-medium text-gray-700">
                  City
                </label>
                <input
                  type="text"
                  name="city"
                  id="city"
                  value={formData.city}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="province" className="block text-sm font-medium text-gray-700">
                  Province/State
                </label>
                <input
                  type="text"
                  name="province"
                  id="province"
                  value={formData.province}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="country" className="block text-sm font-medium text-gray-700">
                  Country
                </label>
                <input
                  type="text"
                  name="country"
                  id="country"
                  value={formData.country}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="postal_code" className="block text-sm font-medium text-gray-700">
                  Postal Code
                </label>
                <input
                  type="text"
                  name="postal_code"
                  id="postal_code"
                  value={formData.postal_code}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
            </div>
          </div>

          {/* Gym Information Section */}
          <div>
            <h3 className="text-xl font-semibold text-gray-900 mb-4">Gym Information</h3>
            <div className="grid grid-cols-1 gap-6 sm:grid-cols-2">
              <div className="sm:col-span-2">
                <label htmlFor="gym_name" className="block text-sm font-medium text-gray-700">
                  Gym Name *
                </label>
                <input
                  type="text"
                  name="gym_name"
                  id="gym_name"
                  required
                  value={formData.gym_name}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div className="sm:col-span-2">
                <label htmlFor="gym_description" className="block text-sm font-medium text-gray-700">
                  Gym Description
                </label>
                <textarea
                  name="gym_description"
                  id="gym_description"
                  rows={3}
                  value={formData.gym_description}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="gym_phone_number" className="block text-sm font-medium text-gray-700">
                  Gym Phone Number
                </label>
                <input
                  type="tel"
                  name="gym_phone_number"
                  id="gym_phone_number"
                  value={formData.gym_phone_number}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div>
                <label htmlFor="gym_email" className="block text-sm font-medium text-gray-700">
                  Gym Email
                </label>
                <input
                  type="email"
                  name="gym_email"
                  id="gym_email"
                  value={formData.gym_email}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
              <div className="sm:col-span-2">
                <label htmlFor="gym_website_url" className="block text-sm font-medium text-gray-700">
                  Gym Website URL
                </label>
                <input
                  type="url"
                  name="gym_website_url"
                  id="gym_website_url"
                  value={formData.gym_website_url}
                  onChange={handleChange}
                  className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                />
              </div>
            </div>
          </div>

          {/* Gym Owner Bio Section */}
          <div>
            <h3 className="text-xl font-semibold text-gray-900 mb-4">About You</h3>
            <div>
              <label htmlFor="brief_bio" className="block text-sm font-medium text-gray-700">
                Brief Bio
              </label>
              <textarea
                name="brief_bio"
                id="brief_bio"
                rows={4}
                value={formData.brief_bio}
                onChange={handleChange}
                className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500"
                placeholder="Tell us about yourself..."
              />
            </div>
          </div>

          {/* Submit Button */}
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
              className="px-6 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {loading ? 'Registering...' : 'Register'}
            </button>
          </div>
        </form>
      </div>
    </div>
  )
}

