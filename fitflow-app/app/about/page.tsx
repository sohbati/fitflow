import Link from 'next/link'

export default function About() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      {/* Navigation */}
      <nav className="bg-white shadow-sm border-b">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between h-16">
            <div className="flex items-center">
              <div className="flex-shrink-0">
                <Link href="/" className="text-2xl font-bold text-blue-600">
                  FitFlow
                </Link>
              </div>
            </div>
            <div className="flex items-center space-x-4">
              <Link 
                href="/" 
                className="text-gray-700 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium"
              >
                Home
              </Link>
              <Link 
                href="/about" 
                className="text-blue-600 px-3 py-2 rounded-md text-sm font-medium font-semibold"
              >
                About
              </Link>
            </div>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <main className="max-w-4xl mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <div className="bg-white rounded-lg shadow-lg p-8">
          <h1 className="text-4xl font-bold text-gray-900 mb-8">About FitFlow</h1>
          
          <div className="prose prose-lg max-w-none">
            <p className="text-gray-600 mb-6">
              FitFlow is a modern Progressive Web App (PWA) designed for fitness tracking and health monitoring. 
              Built with Next.js and optimized for mobile devices, it provides a seamless experience across 
              all platforms.
            </p>

            <h2 className="text-2xl font-semibold text-gray-900 mt-8 mb-4">Key Features</h2>
            <ul className="list-disc list-inside text-gray-600 mb-6 space-y-2">
              <li><strong>Progressive Web App:</strong> Can be installed like a native app with modern web technologies</li>
              <li><strong>Responsive Design:</strong> Optimized for mobile, tablet, and desktop devices</li>
              <li><strong>Google Authentication:</strong> Secure sign-in with Google OAuth 2.0</li>
              <li><strong>Fast Performance:</strong> Built with Next.js for optimal speed and performance</li>
              <li><strong>Modern UI:</strong> Clean, intuitive interface built with Tailwind CSS</li>
            </ul>

            <h2 className="text-2xl font-semibold text-gray-900 mt-8 mb-4">Technology Stack</h2>
            <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
              <div className="bg-gray-50 p-4 rounded-lg">
                <h3 className="font-semibold text-gray-900 mb-2">Frontend</h3>
                <ul className="text-sm text-gray-600 space-y-1">
                  <li>• Next.js 15 (App Router)</li>
                  <li>• TypeScript</li>
                  <li>• Tailwind CSS</li>
                  <li>• React 18</li>
                </ul>
              </div>
              <div className="bg-gray-50 p-4 rounded-lg">
                <h3 className="font-semibold text-gray-900 mb-2">PWA Features</h3>
                <ul className="text-sm text-gray-600 space-y-1">
                  <li>• Service Worker</li>
                  <li>• Web App Manifest</li>
                  <li>• Install Prompts</li>
                  <li>• Push Notifications</li>
                </ul>
              </div>
            </div>

            <h2 className="text-2xl font-semibold text-gray-900 mt-8 mb-4">Mobile App Development</h2>
            <p className="text-gray-600 mb-4">
              FitFlow is designed to be easily converted into native mobile apps using Ionic Capacitor. 
              This allows for:
            </p>
            <ul className="list-disc list-inside text-gray-600 mb-6 space-y-2">
              <li>Cross-platform development (iOS and Android)</li>
              <li>Native device features access</li>
              <li>App store distribution</li>
              <li>Seamless code sharing between web and mobile</li>
            </ul>

            <h2 className="text-2xl font-semibold text-gray-900 mt-8 mb-4">Getting Started</h2>
            <p className="text-gray-600 mb-4">
              To get started with FitFlow:
            </p>
            <ol className="list-decimal list-inside text-gray-600 mb-6 space-y-2">
              <li>Install the app using the "Install App" button on the home page</li>
              <li>Sign in with your Google account to access all features</li>
              <li>For developers, check the README.md for setup instructions</li>
            </ol>

            <div className="mt-8 p-4 bg-blue-50 rounded-lg">
              <p className="text-blue-800 text-sm">
                <strong>Note:</strong> This is a demo PWA showcasing modern web technologies. 
                In a production environment, you would integrate with backend services, 
                user authentication, and data persistence.
              </p>
            </div>
          </div>

          <div className="mt-8 flex justify-center">
            <Link 
              href="/"
              className="inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700"
            >
              Back to Home
            </Link>
          </div>
        </div>
      </main>
    </div>
  )
}
