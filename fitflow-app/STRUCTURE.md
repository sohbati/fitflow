# FitFlow App Structure

## Overview
This document describes the organized folder structure for the FitFlow mobile app, following Next.js 15 best practices.

## Folder Structure

```
fitflow-app/
├── app/                          # Next.js App Router
│   ├── (auth)/                   # Route group for authentication pages
│   │   ├── layout.tsx            # Auth-specific layout (no navigation)
│   │   ├── signin/
│   │   │   └── page.tsx          # Sign in page
│   │   ├── callback/
│   │   │   └── page.tsx          # Google OAuth callback handler
│   │   └── error/
│   │       └── page.tsx          # Authentication error page
│   ├── (main)/                   # Route group for main app pages
│   │   ├── layout.tsx            # Main layout with navigation
│   │   ├── page.tsx              # Root redirect to /home
│   │   ├── home/
│   │   │   └── page.tsx          # Home page (main landing)
│   │   └── about/
│   │       └── page.tsx          # About page
│   ├── api/                      # API routes
│   │   └── auth/
│   │       └── [...nextauth]/
│   │           └── route.ts      # NextAuth API route
│   ├── layout.tsx                # Root layout
│   └── page.tsx                   # Root page (redirects to /home)
│
├── components/                    # Reusable React components
│   ├── auth/                     # Authentication components
│   │   ├── GoogleSignInButton.tsx
│   │   └── AuthStatus.tsx
│   └── layout/                   # Layout components
│       ├── Navigation.tsx
│       └── MobileLayout.tsx
│
├── hooks/                         # Custom React hooks
│   ├── useAuth.ts                # Authentication state management
│   └── useGoogleAuth.ts          # Google OAuth logic
│
├── lib/                           # Utility libraries
│   └── google-auth.ts
│
├── types/                         # TypeScript type definitions
│   └── next-auth.d.ts
│
└── public/                        # Static assets
    ├── manifest.json
    └── icons/
```

## Route Groups

### `(auth)` - Authentication Routes
- **Purpose**: Pages related to authentication (sign in, callbacks, errors)
- **Layout**: Minimal layout without navigation bar
- **Routes**:
  - `/auth/signin` - Sign in page
  - `/auth/callback` - Google OAuth callback
  - `/auth/error` - Authentication error page

### `(main)` - Main Application Routes
- **Purpose**: Main application pages with navigation
- **Layout**: Full layout with navigation bar
- **Routes**:
  - `/home` - Home page (main landing)
  - `/about` - About page
  - `/` - Redirects to `/home`

## Key Features

### 1. Route Groups
- Use parentheses `(auth)` and `(main)` to organize routes without affecting URLs
- Each group has its own layout for consistent styling

### 2. Component Organization
- **Layout Components**: Reusable layout structures
- **Auth Components**: Authentication-related UI components
- **Shared Components**: Can be added as needed

### 3. Custom Hooks
- **useAuth**: Manages authentication state (user, loading, sign out)
- **useGoogleAuth**: Handles Google OAuth flow

### 4. Mobile-First Design
- Responsive layouts optimized for mobile devices
- PWA-ready structure
- Touch-friendly UI components

## Best Practices Followed

1. **Separation of Concerns**: Logic separated from UI components
2. **Reusability**: Components and hooks are reusable across pages
3. **Type Safety**: TypeScript types defined for all data structures
4. **Route Organization**: Route groups for logical grouping
5. **Mobile Optimization**: Mobile-first responsive design

## Adding New Pages

### Adding a Main App Page
1. Create file: `app/(main)/your-page/page.tsx`
2. It will automatically use the main layout with navigation
3. Accessible at `/your-page`

### Adding an Auth Page
1. Create file: `app/(auth)/your-auth-page/page.tsx`
2. It will use the minimal auth layout
3. Accessible at `/auth/your-auth-page`

## Mobile App Integration

This structure is designed to work seamlessly with:
- **Capacitor**: For native mobile app conversion
- **PWA**: Progressive Web App features
- **Responsive Design**: Works on all screen sizes

The organized structure makes it easy to:
- Convert to native mobile apps
- Maintain and scale the codebase
- Add new features and pages
- Test components in isolation

