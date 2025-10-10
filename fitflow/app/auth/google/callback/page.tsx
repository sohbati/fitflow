'use client'

import { useEffect, useState } from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import Link from 'next/link';

export default function GoogleCallbackPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const [status, setStatus] = useState<'loading' | 'success' | 'error'>('loading');
  const [message, setMessage] = useState('');

  useEffect(() => {
    const code = searchParams.get('code');
    const error = searchParams.get('error');

    if (error) {
      setStatus('error');
      setMessage(`Google authentication failed: ${error}`);
      return;
    }

    if (code) {
      // Send the code to your backend IAM service
      const authenticateWithGoogle = async () => {
        try {
          const response = await fetch(`${process.env.NEXT_PUBLIC_IAM_SERVICE_URL}/auth/google`, {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ code }),
          });

          if (!response.ok) {
            const errorData = await response.json();
            throw new Error(errorData.error || 'Google authentication failed');
          }

          const data = await response.json();
          
          // Store token and user data in localStorage for now
          localStorage.setItem('auth_token', data.token);
          localStorage.setItem('user_data', JSON.stringify(data.user));
          
          setStatus('success');
          setMessage('Successfully signed in with Google!');
          
          // Redirect to home page after a short delay
          setTimeout(() => {
            router.push('/');
          }, 2000);
          
        } catch (error: any) {
          console.error('Google authentication error:', error);
          setStatus('error');
          setMessage(`Authentication failed: ${error.message}`);
        }
      };
      
      authenticateWithGoogle();
    } else {
      setStatus('error');
      setMessage('Google authentication failed: No code received.');
    }
  }, [searchParams, router]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <div className="text-center">
          <Link href="/" className="text-3xl font-bold text-blue-600">
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
                  href="/"
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
  );
}
