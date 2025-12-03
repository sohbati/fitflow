import { redirect } from 'next/navigation'

// Redirect root to home page
export default function RootPage() {
  redirect('/home')
}

