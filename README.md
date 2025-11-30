# 👗 Wardrobe Outfit Generator - Complete Application

> A beautiful, full-stack web application for managing your wardrobe and generating stylish outfit combinations

![Status](https://img.shields.io/badge/Status-Production%20Ready-brightgreen)
![Frontend](https://img.shields.io/badge/Frontend-Vanilla%20JS-yellow)
![Backend](https://img.shields.io/badge/Backend-Go-00ADD8)
![Database](https://img.shields.io/badge/Database-PostgreSQL-336791)

---

## 📖 Table of Contents
- [Overview](#overview)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Quick Start](#quick-start)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Frontend Architecture](#frontend-architecture)
- [Design System](#design-system)
- [Deployment](#deployment)
- [Contributing](#contributing)

---

## 🎯 Overview

The Wardrobe Outfit Generator is a comprehensive web application that helps users:
- Digitally organize their clothing collection
- Generate outfit combinations based on preferences
- Plan outfits for different occasions and seasons
- Manage their wardrobe efficiently

**Key Highlights:**
- 🎨 **Premium UI/UX** - Glassmorphism design with smooth animations
- 🔒 **Secure Authentication** - JWT-based auth with token refresh
- 📱 **Fully Responsive** - Works perfectly on all devices
- ⚡ **Fast & Efficient** - Optimized performance
- 🎭 **No Frameworks** - Pure HTML, CSS, JavaScript

---

## ✨ Features

### User Authentication
- ✅ User registration with validation
- ✅ Secure login with JWT tokens
- ✅ Automatic token refresh
- ✅ Protected routes
- ✅ Session management

### Wardrobe Management
- ✅ Add clothing items with images
- ✅ Edit existing items
- ✅ Delete items with confirmation
- ✅ View all clothes in beautiful grid
- ✅ Filter by category (Top, Bottom, Shoes, Accessories, Outerwear)
- ✅ Filter by season (Spring, Summer, Autumn, Winter)
- ✅ Image support via URLs

### Outfit Generation
- ✅ AI-powered outfit combinations
- ✅ Filter by occasion and season
- ✅ View generated outfits
- ✅ Detailed outfit view
- ✅ Delete unwanted outfits
- ✅ Browse recent vs all outfits

### UI/UX
- ✅ Premium glassmorphism design
- ✅ Smooth animations and transitions
- ✅ Loading states
- ✅ Error handling
- ✅ Empty states
- ✅ Form validation
- ✅ Mobile-responsive
- ✅ Touch-friendly

---

## 🛠️ Technology Stack

### Frontend
- **HTML5** - Semantic markup
- **CSS3** - Glassmorphism, Grid, Flexbox, Animations
- **JavaScript (ES6+)** - Modules, Async/Await, Fetch API

### Backend
- **Go** - High-performance server
- **Gin** - Web framework
- **PostgreSQL** - Database
- **JWT** - Authentication
- **bcrypt** - Password hashing

### Tools & Libraries
- **Viper** - Configuration management
- **godotenv** - Environment variables
- **logrus** - Logging

---

## 🚀 Quick Start

### Prerequisites
- Go 1.16+ installed
- PostgreSQL database running
- Modern web browser

### Installation

1. **Clone the repository**
```bash
git clone <repository-url>
cd SalimProject
```

2. **Configure environment**
Create `.env` file:
```env
DB_PASSWORD=your_database_password
```

3. **Update config**
Edit `configs/config.yml` if needed:
```yaml
port: "8080"
db:
  host: "localhost"
  port: "5432"
  username: "postgres"
  dbname: "wardrobe_db"
  sslmode: "disable"
```

4. **Run database migrations**
```bash
# Your migration commands here
```

5. **Start the application**
```bash
# Option 1: Using the batch file (Windows)
start.bat

# Option 2: Direct Go command
go run cmd/main.go
```

6. **Access the application**
Open your browser and navigate to:
```
http://localhost:8080
```

---

## 📁 Project Structure

```
SalimProject/
├── cmd/
│   └── main.go                 # Application entry point
├── configs/
│   └── config.yml              # Configuration file
├── migrations/
│   ├── 000001_create_users_table.up.sql
│   ├── 000002_create_clothes_table.up.sql
│   └── 000003_create_outfits_table.up.sql
├── models/
│   ├── user.go
│   ├── clothes.go
│   └── outfit.go
├── pkg/
│   ├── dto/
│   │   └── dto.go              # Data transfer objects
│   ├── handler/
│   │   ├── handler.go          # Router & static files
│   │   ├── auth.go             # Auth handlers
│   │   ├── clothes.go          # Clothes handlers
│   │   ├── outfit.go           # Outfit handlers
│   │   ├── middleware.go       # JWT middleware
│   │   └── response.go         # Response helpers
│   ├── repository/
│   │   ├── repository.go
│   │   ├── db.go
│   │   ├── auth.go
│   │   ├── clothes.go
│   │   └── outfit.go
│   └── service/
│       ├── service.go
│       ├── auth.go
│       ├── clothes.go
│       └── outfit.go
├── server/
│   └── server.go               # HTTP server
├── public/                     # Frontend files
│   ├── index.html
│   ├── login.html
│   ├── signup.html
│   ├── wardrobe.html
│   ├── outfit.html
│   ├── css/
│   │   ├── styles.css
│   │   └── pages/
│   │       ├── auth.css
│   │       ├── home.css
│   │       ├── wardrobe.css
│   │       └── outfit.css
│   └── js/
│       ├── config.js
│       ├── api/
│       │   ├── auth.js
│       │   ├── clothes.js
│       │   └── outfit.js
│       ├── helpers/
│       │   └── auth.js
│       └── pages/
│           ├── login.js
│           ├── signup.js
│           ├── home.js
│           ├── wardrobe.js
│           └── outfit.js
├── .env
├── go.mod
├── go.sum
├── Dockerfile
├── docker-compose.yml
└── README.md
```

---

## 📡 API Documentation

### Base URL
```
http://localhost:8080
```

### Authentication Endpoints

#### Register User
```http
POST /auth/sign-up
Content-Type: application/json

{
  "name": "John Doe",
  "username": "johndoe",
  "password": "password123"
}
```

#### Login
```http
POST /auth/sign-in
Content-Type: application/json

{
  "username": "johndoe",
  "password": "password123"
}

Response:
{
  "accessToken": "eyJhbGc...",
  "refreshToken": "eyJhbGc..."
}
```

#### Refresh Token
```http
POST /auth/refresh
Content-Type: application/json

{
  "refreshToken": "eyJhbGc..."
}
```

### Clothes Endpoints

All clothes endpoints require authentication:
```
Authorization: Bearer {accessToken}
```

#### List All Clothes
```http
GET /api/clothes/item
```

#### Get Clothes by ID
```http
GET /api/clothes/item/{id}
```

#### Add Clothes
```http
POST /api/clothes/item
Content-Type: application/json

{
  "name": "Blue Denim Jacket",
  "category": "outerwear",
  "color": "Blue",
  "season": "autumn",
  "material": "Denim",
  "image_url": "https://example.com/image.jpg"
}
```

#### Update Clothes
```http
PUT /api/clothes/item/{id}
Content-Type: application/json

{
  "name": "Updated Name",
  "category": "top",
  "color": "Red",
  "season": "summer",
  "material": "Cotton",
  "image_url": "https://example.com/new.jpg"
}
```

#### Delete Clothes
```http
DELETE /api/clothes/item/{id}
```

### Outfit Endpoints

#### Generate Outfit
```http
POST /api/outfit/generate
Authorization: Bearer {accessToken}
Content-Type: application/json

{
  "occasion": "Business",
  "season": "autumn"
}
```

#### List All Outfits
```http
GET /api/outfit/
Authorization: Bearer {accessToken}
```

#### Get Outfit by ID
```http
GET /api/outfit/{id}
Authorization: Bearer {accessToken}
```

#### Delete Outfit
```http
DELETE /api/outfit/{id}
Authorization: Bearer {accessToken}
```

---

## 🎨 Frontend Architecture

### Design Principles
1. **Separation of Concerns** - API, Helpers, Pages separated
2. **Modularity** - ES6 modules for clean imports
3. **Reusability** - Shared components and utilities
4. **No Inline Code** - All CSS and JS in separate files
5. **Progressive Enhancement** - Works without JS for basic content

### File Organization
```
public/
├── HTML Pages (5 files)
├── CSS Styles (5 files)
│   ├── Global styles
│   └── Page-specific styles
└── JavaScript (10 files)
    ├── Configuration
    ├── API layer
    ├── Helper functions
    └── Page controllers
```

### State Management
- **localStorage** for JWT tokens
- **In-memory** for page state
- **API calls** for data fetching

### Error Handling
- Try-catch blocks around all API calls
- User-friendly error messages
- Loading states during operations
- Empty states when no data

---

## 🎨 Design System

### Color Palette
```css
Primary:   #8B7FFF (Purple)
Secondary: #FF9ECD (Pink)
Accent:    #FFD17A (Yellow)
Success:   #7FD1AE (Green)
Danger:    #FF8B94 (Red)
```

### Typography
```css
Font: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif
Sizes: 0.75rem to 3rem
Weights: 400, 600, 700
```

### Components
- Buttons (Primary, Secondary, Outline, Danger)
- Forms (Inputs, Selects, Validation)
- Cards (Standard, Glassmorphism)
- Modals (Centered, Animated)
- Loading States (Spinners, Skeletons)

### Animations
- Fade in/out
- Slide up/down
- Scale on hover
- Continuous float
- Smooth transitions (150ms - 350ms)

---

## 🚢 Deployment

### Production Build

1. **Build Go application**
```bash
go build -o wardrobe-app cmd/main.go
```

2. **Set environment variables**
```bash
export DB_PASSWORD=your_production_password
```

3. **Run application**
```bash
./wardrobe-app
```

### Docker Deployment

```bash
docker-compose up -d
```

### Nginx Configuration (Optional)
```nginx
server {
    listen 80;
    server_name yourdomain.com;
    
    location / {
        proxy_pass http://localhost:8080;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
    }
}
```

---

## 📚 Additional Documentation

- **[SETUP_GUIDE.md](SETUP_GUIDE.md)** - Detailed setup and testing instructions
- **[DESIGN_PREVIEW.md](DESIGN_PREVIEW.md)** - UI/UX design documentation
- **[FRONTEND_COMPLETE.md](FRONTEND_COMPLETE.md)** - Frontend implementation summary
- **[public/README.md](public/README.md)** - Frontend-specific documentation

---

## 🧪 Testing

### Manual Testing Checklist

1. **Authentication**
   - [ ] Register new user
   - [ ] Login with credentials
   - [ ] Token refresh works
   - [ ] Logout functionality
   - [ ] Protected routes redirect

2. **Wardrobe**
   - [ ] Add new clothing item
   - [ ] Edit existing item
   - [ ] Delete item
   - [ ] Filter by category
   - [ ] Filter by season
   - [ ] Images display correctly

3. **Outfits**
   - [ ] Generate outfit
   - [ ] View outfit details
   - [ ] Delete outfit
   - [ ] Filter recent/all

4. **Responsive**
   - [ ] Desktop (1920px)
   - [ ] Tablet (768px)
   - [ ] Mobile (375px)

---

## 🐛 Troubleshooting

### Common Issues

**Backend won't start:**
- Check PostgreSQL is running
- Verify database credentials in `.env`
- Ensure migrations are applied

**Frontend can't connect:**
- Verify backend is running on port 8080
- Check browser console for errors
- Verify CORS is enabled

**Images not loading:**
- Use direct image URLs
- Check image URL is publicly accessible
- Try Unsplash images for testing

---

## 📝 License

This project is part of the SalimProject educational application.

---

## 👨‍💻 Author

Created for SalimProject

---

## 🙏 Acknowledgments

- Design inspiration from modern fashion apps
- Icons using Unicode emojis
- CSS techniques: Glassmorphism, Neumorphism
- Images from Unsplash (for testing)

---

## 📞 Support

For issues or questions:
1. Check the [SETUP_GUIDE.md](SETUP_GUIDE.md)
2. Review browser console for errors
3. Check backend logs
4. Verify database connection

---

**⭐ Star this repository if you find it helpful!**

Made with ❤️ and lots of ☕

