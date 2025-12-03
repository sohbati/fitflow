'use client'

import Link from 'next/link'
import { GoogleSignInButton } from '@/components/auth/GoogleSignInButton'

export default function SignInPage() {
  return (
    <div className="flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div>
          <div className="flex justify-center">
            <Link href="/home" className="text-3xl font-bold text-blue-600">
              FitFlow
            </Link>
          </div>
          <h2 className="mt-6 text-center text-3xl font-extrabold text-gray-900">
            Sign in to your account
          </h2>
          <p className="mt-2 text-center text-sm text-gray-600">
            Welcome back! Please sign in to continue using FitFlow.
          </p>
        </div>
        
        <div className="mt-8 space-y-6">
          <GoogleSignInButton fullWidth variant="default" />
          
          <div className="text-center">
            <Link 
              href="/home"
              className="text-blue-600 hover:text-blue-800 text-sm underline"
            >
              Back to home
            </Link>
          </div>
        </div>
      </div>
    </div>
  )
}

