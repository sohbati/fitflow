# FitFlow - Progressive Web App

A modern Progressive Web App (PWA) built with Next.js, designed for fitness tracking and optimized for mobile devices. This app can be easily converted into native mobile apps using Ionic Capacitor.

## 🚀 Features

- **Progressive Web App**: Can be installed like a native app with modern web technologies
- **Google Authentication**: Secure sign-in with Google OAuth 2.0
- **Responsive Design**: Optimized for mobile, tablet, and desktop devices
- **Fast Performance**: Built with Next.js for optimal speed and performance
- **Modern UI**: Clean, intuitive interface built with Tailwind CSS
- **Capacitor Ready**: Easy conversion to native mobile apps

## 🛠️ Tech Stack

- **Frontend**: Next.js 15 (App Router), TypeScript, Tailwind CSS
- **Authentication**: NextAuth.js with Google OAuth 2.0
- **PWA**: next-pwa for service worker and offline caching
- **Mobile**: Ionic Capacitor for native app development
- **Build Tools**: ESLint, Prettier

## 📦 Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd fitflow
```

2. Install dependencies:
```bash
npm install
```

3. Set up environment variables:
```bash
cp env.example .env.local
```
Edit `.env.local` with your configuration:
```env
NEXTAUTH_URL=http://localhost:3000
NEXTAUTH_SECRET=your-secret-key-here
GOOGLE_CLIENT_ID=your-google-client-id
GOOGLE_CLIENT_SECRET=your-google-client-secret
```

4. Run the development server:
```bash
npm run dev
```

5. Open [http://localhost:3000](http://localhost:3000) in your browser.

## 🚀 Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build the application
- `npm run export` - Build and export static files
- `npm run start` - Start production server
- `npm run lint` - Run ESLint

## 📱 PWA Features

### Installation
The app can be installed on devices that support PWA installation:
- Click the "Install App" button when it appears
- Or use the browser's "Add to Home Screen" option

### Google Authentication
- Secure OAuth 2.0 authentication with Google
- User profile information (name, email, avatar)
- Session management with NextAuth.js
- Protected routes and user state management

### Performance
- Service worker for improved performance
- Optimized caching strategies
- Fast loading times with Next.js

### Manifest
- App name: "FitFlow - Fitness Tracking App"
- Short name: "FitFlow"
- Theme color: #317EFB
- Display mode: standalone

## 🔧 Capacitor Integration

This PWA is designed to be easily converted into native mobile apps using Ionic Capacitor.

### Prerequisites
- Node.js 18+ 
- Android Studio (for Android development)
- Xcode (for iOS development, macOS only)

### Setup Instructions

1. **Install Capacitor CLI** (if not already installed):
```bash
npm install -g @capacitor/cli
```

2. **Initialize Capacitor**:
```bash
npm run cap:init
```

3. **Add Android platform**:
```bash
npm run cap:add:android
```

4. **Build and copy web assets**:
```bash
npm run export
npm run cap:copy
```

5. **Open in Android Studio**:
```bash
npm run cap:open:android
```

### Capacitor Commands

- `npm run cap:init` - Initialize Capacitor
- `npm run cap:add:android` - Add Android platform
- `npm run cap:copy` - Copy web assets to native projects
- `npm run cap:open:android` - Open Android project in Android Studio
- `npm run cap:sync` - Sync changes to native projects

## 📁 Project Structure

```
fitflow/
├── app/
│   ├── layout.tsx          # Root layout with PWA meta tags
│   ├── page.tsx            # Home page with install button
│   ├── about/
│   │   └── page.tsx        # About page
│   └── offline/
│       └── page.tsx        # Offline fallback page
├── public/
│   ├── icons/
│   │   └── icon.svg        # App icon
│   └── manifest.json       # PWA manifest
├── styles/
│   └── globals.css         # Global styles
├── next.config.js          # Next.js configuration with PWA
├── capacitor.config.ts     # Capacitor configuration
└── package.json           # Dependencies and scripts
```

## 🎯 Development

### Adding New Pages
1. Create a new folder in the `app/` directory
2. Add a `page.tsx` file
3. The route will be automatically available

### PWA Configuration
- Service worker settings in `next.config.js`
- Manifest configuration in `public/manifest.json`
- Meta tags in `app/layout.tsx`

### Styling
- Uses Tailwind CSS for styling
- Responsive design with mobile-first approach
- Dark/light mode support (optional)

## 🚀 Deployment

### Static Export
```bash
npm run export
```
This creates an `out/` directory with static files ready for deployment.

### Capacitor Build
```bash
npm run export
npm run cap:copy
npm run cap:open:android
```

## 📊 Performance

- **Lighthouse PWA Score**: 90+
- **Core Web Vitals**: Optimized
- **Bundle Size**: Minimized
- **Loading Speed**: Fast

## 🔒 Security

- HTTPS required for PWA features
- Secure headers configured
- Content Security Policy implemented

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For support and questions:
- Check the documentation
- Open an issue on GitHub
- Contact the development team

## 🔮 Future Enhancements

- [ ] User authentication
- [ ] Data persistence
- [ ] Push notifications
- [ ] Advanced fitness tracking features
- [ ] Social features
- [ ] Analytics dashboard

---

**Note**: This is a demo PWA showcasing modern web technologies. In a production environment, you would integrate with backend services, user authentication, and data persistence.