'use client'

import { ReactNode } from 'react'
import { Navigation } from './Navigation'

interface MobileLayoutProps {
  children: ReactNode
}

export function MobileLayout({ children }: MobileLayoutProps) {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <Navigation />
      {/* Add top padding to account for fixed navbar */}
      <main className="max-w-7xl mx-auto pt-20 pb-12 px-4 sm:px-6 lg:px-8">
        {children}
      </main>
    </div>
  )
}

