# CoachAidApp - Project Planning Summary

## Project Overview
**Goal**: Develop an innovative fitness/coaching app as a solo developer project for free time development, with potential to grow into a sustainable business.

**Current Status**: Existing codebase from previous year, now restarting with modern approach and clear MVP focus.

## MVP Requirements

### Core User Roles
1. **Coach**
   - Register/Login
   - Create workout plans
   - Track clients
   - View dashboard with progress analytics

2. **Trainer**
   - Register/Login
   - View assigned clients/workouts
   - Mark sessions complete
   - Basic stats and progress tracking

### Public Features (No Login Required)
- Landing page with app overview/benefits
- Sample workouts (demo of AI-generated or pre-made plans)
- Blog/tips section (health & fitness content)
- Contact/request demo form
- Optional: "Try as guest" mode for interactive sample workout

### Exercise Location Support
Four location types for workout customization:
1. **Home** - Equipment-limited workouts
2. **Gym** - Full equipment access
3. **Park** - Outdoor bodyweight exercises
4. **Nature** - Mountain/hiking/outdoor activities

### User Registration Context
- **Gym Registration Flag**: Both coaches and trainers can indicate if they're registered at a specific gym
- **Gym Information**: Optional gym name and location storage
- **Location-based Filtering**: Enable finding local coaches/trainers

## Technical Architecture

### Project Structure
```
CoachAidApp/
├── backend/
│   ├── src/
│   │   ├── controllers/     # API endpoints
│   │   ├── models/          # DB schemas/ORM models
│   │   ├── services/        # Business logic
│   │   ├── utils/           # Helper functions
│   │   └── config/          # DB, env, JWT configs
│   ├── tests/               # Unit/integration tests
│   └── package.json/pom.xml
├── frontend/
│   ├── src/
│   │   ├── components/      # Reusable UI components
│   │   ├── pages/           # Views for Coach, Trainer, Public
│   │   ├── services/        # API calls
│   │   ├── context/         # State management
│   │   ├── routes/          # Routing configuration
│   │   └── utils/           # Helpers (formatting, validation)
│   ├── public/              # Assets
│   └── package.json
└── README.md
```

### Database Schema
**User Table**
```sql
User {
    id: UUID
    name: string
    email: string
    password_hash: string
    role: ENUM('coach','trainer')
    gym_registered: boolean
    gym_name: string (nullable)
    gym_location: string (nullable)
}
```

**Workout Table**
```sql
Workout {
    id: UUID
    title: string
    description: text
    exercises: JSON/relation table
    location: ENUM('home','gym','park','nature')
    coach_id: UUID
    created_at: timestamp
    updated_at: timestamp
}
```

**Exercise Table** (Optional for expansion)
```sql
Exercise {
    id: UUID
    name: string
    equipment: string
    default_location: ENUM('home','gym','park','nature')
    instructions: text
    difficulty_level: ENUM('beginner','intermediate','advanced')
}
```

## MVP Views/Pages

### Public Views
- **Home/Landing Page**: App overview, benefits, call-to-action
- **About/Features**: Detailed feature descriptions
- **Blog/Tips**: Health & fitness content
- **Contact/Request Demo**: Lead generation form
- **Sample Workouts**: Location-based workout demos

### Coach Views
- **Login/Register**: Authentication
- **Dashboard**: Clients summary, recent activity
- **Client Management**: List, details, progress tracking
- **Workout Creation**: Location-aware exercise planning
- **Analytics**: Progress reports, client statistics

### Trainer Views
- **Login/Register**: Authentication
- **Assigned Clients**: List of clients and their workouts
- **Session Tracking**: Mark workouts complete, add notes
- **Progress View**: Basic stats and client progress

## Technology Stack Recommendations

### Frontend
- **Framework**: React + TailwindCSS (or React Native for mobile)
- **State Management**: React Context + hooks (lightweight)
- **Routing**: React Router
- **UI Components**: Headless UI or custom components

### Backend
- **Runtime**: Node.js (Express) or Python FastAPI
- **Database**: PostgreSQL (production) or SQLite (MVP)
- **Authentication**: JWT + bcrypt
- **Validation**: Input validation middleware

### Deployment
- **Frontend**: Vercel/Netlify
- **Backend**: Render/Railway
- **Database**: PostgreSQL (managed service)

### Optional AI Integration
- **OpenAI API**: For personalized workout plan generation
- **Future**: Pose estimation for form feedback

## Development Phases

### Phase 1: Core MVP (2-3 months)
1. Set up project structure
2. Implement authentication (Coach/Trainer roles)
3. Create basic workout CRUD operations
4. Add location-based exercise filtering
5. Build essential views (dashboard, workout creation)

### Phase 2: Enhancement (1-2 months)
1. Add gym registration features
2. Implement client-trainer assignment
3. Create progress tracking
4. Add public features and landing page

### Phase 3: AI Integration (1-2 months)
1. Integrate OpenAI for workout personalization
2. Add intelligent exercise recommendations
3. Implement location-aware workout generation

## Business Model Considerations

### Monetization Strategies
1. **Freemium Model**: Basic features free, premium AI coaching paid
2. **Subscription Tiers**: Coach plans, Trainer plans, Premium features
3. **B2B Opportunities**: Gym management features for fitness facilities
4. **Affiliate Marketing**: Fitness equipment and supplement partnerships

### Growth Opportunities
1. **Community Features**: User-generated content, challenges
2. **Mobile App**: React Native for iOS/Android
3. **Wearable Integration**: Apple Watch, Fitbit data sync
4. **Local Networking**: Connect coaches/trainers with clients by location

## Next Steps

### Immediate Actions
1. **Set up development environment** with chosen tech stack
2. **Create database schema** and initial migrations
3. **Implement authentication system** with role-based access
4. **Build basic workout CRUD** with location support
5. **Create MVP UI** for core user flows

### Success Metrics
- User registration and retention rates
- Workout creation and completion rates
- User engagement with location-based features
- Coach-trainer collaboration effectiveness

## Notes
- Focus on solo development approach - avoid complex team dependencies
- Prioritize automation and tech leverage over manual processes
- Build for scalability but start with simple, working features
- Consider open-source components to accelerate development
- Plan for future AI integration from the beginning

---
*This summary is based on ChatGPT discussion analysis and provides a roadmap for developing CoachAidApp as an innovative fitness coaching platform.*
