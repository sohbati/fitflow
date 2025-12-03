'use client'

export async function initiateGoogleSignIn(): Promise<void> {
  const iamServiceUrl = process.env.NEXT_PUBLIC_IAM_SERVICE_URL
  
  if (!iamServiceUrl) {
    throw new Error('Configuration error: IAM service URL is not set. Please check your .env.local file and restart the dev server.')
  }
  
  console.log('Fetching Google auth URL from:', `${iamServiceUrl}/auth/google/url`)
  
  const response = await fetch(`${iamServiceUrl}/auth/google/url`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    mode: 'cors',
  })
  
  console.log('Response status:', response.status, response.statusText)
  
  if (!response.ok) {
    const errorText = await response.text()
    console.error('Response error:', errorText)
    throw new Error(`Failed to get Google auth URL: ${response.status} ${response.statusText}`)
  }
  
  const data = await response.json()
  console.log('Received auth URL data:', data)
  
  if (!data.url) {
    throw new Error('Invalid response: no URL in response')
  }
  
  // Redirect to Google's auth page
  window.location.href = data.url
}

export async function handleGoogleCallback(code: string): Promise<{ token: string; user: any }> {
  const iamServiceUrl = process.env.NEXT_PUBLIC_IAM_SERVICE_URL
  
  if (!iamServiceUrl) {
    throw new Error('IAM service URL is not configured')
  }

  const response = await fetch(`${iamServiceUrl}/auth/google`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ code }),
  })

  if (!response.ok) {
    const errorData = await response.json().catch(() => ({}))
    throw new Error(errorData.error || 'Google authentication failed')
  }

  const data = await response.json()
  
  // Store token and user data
  localStorage.setItem('auth_token', data.token)
  localStorage.setItem('user_data', JSON.stringify(data.user))
  
  return data
}
