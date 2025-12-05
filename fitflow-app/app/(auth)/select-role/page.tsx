'use client'

import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { useState } from 'react'
import { Navigation } from '@/components/layout/Navigation'

export default function SelectRolePage() {
  const router = useRouter()
  const [selectedRole, setSelectedRole] = useState<string | null>(null)

  const handleRoleSelect = (role: string) => {
    setSelectedRole(role)
    
    // Redirect based on selected role
    switch (role) {
      case 'gym-owner':
        router.push('/register/gym-owner')
        break
      case 'trainer':
        router.push('/register/trainer')
        break
      case 'trainee':
        router.push('/register/trainee')
        break
      default:
        break
    }
  }

  return (
    <>
      <Navigation />
      <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 pt-24">
        <div className="max-w-2xl w-full space-y-8">
        <div className="text-center">
          <Link href="/home" className="text-3xl font-bold text-blue-600">
            FitFlow
          </Link>
          <h2 className="mt-6 text-3xl font-extrabold text-gray-900">
            Welcome to FitFlow!
          </h2>
          <p className="mt-2 text-sm text-gray-600">
            Please select your role to continue
          </p>
        </div>

        <div className="bg-white rounded-lg shadow-md p-8">
          <div className="grid grid-cols-1 gap-6 sm:grid-cols-3">
            {/* Gym Owner Card */}
            <button
              onClick={() => handleRoleSelect('gym-owner')}
              className={`relative rounded-lg border-2 p-6 text-left transition-all hover:shadow-lg ${
                selectedRole === 'gym-owner'
                  ? 'border-blue-600 bg-blue-50'
                  : 'border-gray-200 hover:border-blue-300'
              }`}
            >
              <div className="flex flex-col items-center text-center">
                <div className="mb-4 rounded-full bg-blue-100 p-4">
                  <svg
                    className="h-8 w-8 text-blue-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"
                    />
                  </svg>
                </div>
                <h3 className="text-lg font-semibold text-gray-900">Gym Owner</h3>
                <p className="mt-2 text-sm text-gray-500">
                  Manage your gym and trainers
                </p>
              </div>
            </button>

            {/* Trainer Card */}
            <button
              onClick={() => handleRoleSelect('trainer')}
              className={`relative rounded-lg border-2 p-6 text-left transition-all hover:shadow-lg ${
                selectedRole === 'trainer'
                  ? 'border-blue-600 bg-blue-50'
                  : 'border-gray-200 hover:border-blue-300'
              }`}
            >
              <div className="flex flex-col items-center text-center">
                <div className="mb-4 rounded-full bg-green-100 p-4">
                  <svg
                    className="h-8 w-8 text-green-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                    />
                  </svg>
                </div>
                <h3 className="text-lg font-semibold text-gray-900">Trainer</h3>
                <p className="mt-2 text-sm text-gray-500">
                  Train clients and manage sessions
                </p>
              </div>
            </button>

            {/* Trainee Card */}
            <button
              onClick={() => handleRoleSelect('trainee')}
              className={`relative rounded-lg border-2 p-6 text-left transition-all hover:shadow-lg ${
                selectedRole === 'trainee'
                  ? 'border-blue-600 bg-blue-50'
                  : 'border-gray-200 hover:border-blue-300'
              }`}
            >
              <div className="flex flex-col items-center text-center">
                <div className="mb-4 rounded-full bg-purple-100 p-4">
                  <svg
                    className="h-8 w-8 text-purple-600"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      strokeWidth={2}
                      d="M13 10V3L4 14h7v7l9-11h-7z"
                    />
                  </svg>
                </div>
                <h3 className="text-lg font-semibold text-gray-900">Trainee</h3>
                <p className="mt-2 text-sm text-gray-500">
                  Join a gym and start training
                </p>
              </div>
            </button>
          </div>
        </div>
        </div>
      </div>
    </>
  )
}

