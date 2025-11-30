# Wardrobe Outfit Generator - Complete Setup Guide

## 📋 Table of Contents
1. [Project Overview](#project-overview)
2. [Quick Start](#quick-start)
3. [Backend API Endpoints](#backend-api-endpoints)
4. [Frontend Pages](#frontend-pages)
5. [Testing the Application](#testing-the-application)
6. [Common Issues](#common-issues)

## 🎯 Project Overview

A full-stack web application for managing your wardrobe and generating outfit combinations with a beautiful, modern UI.

**Tech Stack:**
- **Backend**: Go, Gin, PostgreSQL
- **Frontend**: Pure HTML, CSS, JavaScript (no frameworks)
- **Authentication**: JWT tokens
- **Design**: Glassmorphism, responsive, mobile-first

## 🚀 Quick Start

### Step 1: Start the Backend Server

```bash
# From the project root directory
cd C:\Users\DWA\GolandProjects\SalimProject

# Make sure PostgreSQL is running and configured in .env

# Run the application
go run cmd/main.go
```

The server will start on `http://localhost:8080`

### Step 2: Access the Application

Open your browser and navigate to:
```
http://localhost:8080
```

You should see the beautiful home page!

### Step 3: Create an Account

1. Click on the "Sign up" link or navigate to `http://localhost:8080/signup.html`
2. Enter your details:
   - Full Name
   - Username (min 3 characters)
   - Password (min 6 characters)
3. Click "Create Account"
4. You'll be redirected to the login page

### Step 4: Login

1. Enter your username and password
2. Click "Sign In"
3. You'll be redirected to the home page

### Step 5: Add Clothes to Your Wardrobe

1. Click "Browse Wardrobe" or navigate to the "My Closet" page
2. Click "Add Item"
3. Fill in the form:
   - **Name**: e.g., "Blue Denim Jacket"
   - **Category**: Choose from dropdown (Top, Bottom, Shoes, Accessories, Outerwear)
   - **Color**: e.g., "Blue"
   - **Season**: Choose from dropdown (Spring, Summer, Autumn, Winter, All Seasons)
   - **Material**: e.g., "Denim"
   - **Image URL**: Paste a direct link to an image
     - Example: `https://images.unsplash.com/photo-1551028719-00167b16eac5?w=500`
4. Click "Save Item"

**Tip**: Use image URLs from these sources:
- Unsplash: https://unsplash.com
- Imgur: https://imgur.com
- Any direct image URL ending in .jpg, .png, .webp

### Step 6: Generate Outfits

1. Navigate to the "Outfits" page
2. Optionally fill in:
   - **Occasion**: e.g., "Business", "Casual", "Party"
   - **Season**: Select preferred season
3. Click "Generate Outfit"
4. View your generated outfit combination!

## 🔌 Backend API Endpoints

### Authentication Endpoints

#### Sign Up
```http
POST /auth/sign-up
Content-Type: application/json

{
  "name": "John Doe",
  "username": "johndoe",
  "password": "password123"
}
```

#### Sign In
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

#### Get All Clothes
```http
GET /api/clothes/item
Authorization: Bearer {accessToken}
```

#### Get Clothes by ID
```http
GET /api/clothes/item/{id}
Authorization: Bearer {accessToken}
```

#### Add Clothes
```http
POST /api/clothes/item
Authorization: Bearer {accessToken}
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
Authorization: Bearer {accessToken}
Content-Type: application/json

{
  "name": "Updated Name",
  "category": "top",
  "color": "Red",
  "season": "summer",
  "material": "Cotton",
  "image_url": "https://example.com/new-image.jpg"
}
```

#### Delete Clothes
```http
DELETE /api/clothes/item/{id}
Authorization: Bearer {accessToken}
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

#### Get All Outfits
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

## 🎨 Frontend Pages

### 1. Home Page (`index.html`)
- **URL**: `http://localhost:8080/` or `http://localhost:8080/index.html`
- **Features**: Hero section, feature cards, navigation
- **Protected**: Yes (requires login)

### 2. Sign Up Page (`signup.html`)
- **URL**: `http://localhost:8080/signup.html`
- **Features**: User registration form
- **Protected**: No

### 3. Login Page (`login.html`)
- **URL**: `http://localhost:8080/login.html`
- **Features**: User authentication form
- **Protected**: No

### 4. Wardrobe Page (`wardrobe.html`)
- **URL**: `http://localhost:8080/wardrobe.html`
- **Features**: 
  - View all clothes in grid layout
  - Add new clothes items
  - Edit existing items
  - Delete items
  - Filter by category and season
- **Protected**: Yes (requires login)

### 5. Outfit Page (`outfit.html`)
- **URL**: `http://localhost:8080/outfit.html`
- **Features**:
  - Generate new outfits
  - View all generated outfits
  - View outfit details
  - Delete outfits
  - Filter by recent/all
- **Protected**: Yes (requires login)

## 🧪 Testing the Application

### Test Scenario 1: Complete User Flow

1. **Register a new user**
   - Go to signup page
   - Enter: Name: "Test User", Username: "testuser", Password: "test123"
   - Verify redirect to login page

2. **Login**
   - Enter credentials
   - Verify redirect to home page
   - Check that navbar shows logout button

3. **Add clothes items** (Add at least 4-5 items for best results)
   - Navigate to wardrobe page
   - Add a Top: "White T-Shirt", category: "top", color: "White"
   - Add a Bottom: "Blue Jeans", category: "bottom", color: "Blue"
   - Add Shoes: "White Sneakers", category: "shoes", color: "White"
   - Add Accessories: "Sunglasses", category: "accessories", color: "Black"

4. **Filter clothes**
   - Use category filter to show only "tops"
   - Use season filter to show summer items

5. **Edit a clothing item**
   - Click edit on any item
   - Change the color or name
   - Verify changes are saved

6. **Generate outfit**
   - Go to outfit page
   - Enter occasion: "Casual"
   - Select season: "Summer"
   - Click generate
   - Verify outfit is created with your clothes

7. **View outfit details**
   - Click "View Details" on generated outfit
   - Verify all items are shown

8. **Delete outfit**
   - Click delete button
   - Confirm deletion
   - Verify outfit is removed

9. **Logout**
   - Click logout button
   - Verify redirect to login page
   - Try accessing wardrobe page - should redirect to login

### Test Scenario 2: Image URLs to Use

Here are some free image URLs you can use for testing:

**Tops:**
- `https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=500`
- `https://images.unsplash.com/photo-1618354691373-d851c5c3a990?w=500`

**Bottoms:**
- `https://images.unsplash.com/photo-1542272454315-7f6fabf03775?w=500`
- `https://images.unsplash.com/photo-1624378439575-d8705ad7ae80?w=500`

**Shoes:**
- `https://images.unsplash.com/photo-1549298916-b41d501d3772?w=500`
- `https://images.unsplash.com/photo-1460353581641-37baddab0fa2?w=500`

**Accessories:**
- `https://images.unsplash.com/photo-1511499767150-a48a237f0083?w=500`
- `https://images.unsplash.com/photo-1576566588028-4147f3842f27?w=500`

## 🐛 Common Issues

### Issue 1: "Failed to fetch" error

**Problem**: API calls are failing with network error

**Solutions**:
1. Verify backend server is running: `go run cmd/main.go`
2. Check database connection in `.env` file
3. Verify PostgreSQL is running
4. Check browser console for detailed error
5. Ensure CORS is enabled (already configured in handler.go)

### Issue 2: Images not displaying

**Problem**: Clothes images show placeholder instead of actual image

**Solutions**:
1. Use direct image URLs (ending in .jpg, .png, .webp)
2. Ensure image URL is publicly accessible
3. Check browser console for CORS errors on image
4. Try using images from Unsplash (CORS-friendly)

### Issue 3: Token expiration

**Problem**: Automatically logged out after some time

**Solutions**:
- This is normal behavior when access token expires
- The app will automatically try to refresh the token
- If refresh fails, you need to login again
- Check backend token expiration settings

### Issue 4: Can't generate outfits

**Problem**: "Failed to generate outfit" error

**Solutions**:
1. Add more clothes to your wardrobe (at least 3-4 items)
2. Ensure you have clothes in different categories
3. Check backend logs for specific errors
4. Try without specifying occasion/season

### Issue 5: Page redirects to login

**Problem**: Accessing wardrobe/outfit page redirects to login

**Solutions**:
1. Make sure you're logged in
2. Check if tokens are stored in localStorage (F12 > Application > Local Storage)
3. Clear localStorage and login again
4. Check if backend is returning proper tokens on login

## 📱 Mobile Testing

The application is fully responsive. Test on mobile by:

1. **Chrome DevTools**:
   - Press F12
   - Click device toolbar icon (Ctrl+Shift+M)
   - Select mobile device
   - Test all features

2. **Actual Mobile Device**:
   - Connect your phone to same network
   - Find your computer's IP address
   - Access `http://YOUR_IP:8080` from mobile browser

## 🎯 Feature Checklist

- ✅ User Registration
- ✅ User Login
- ✅ JWT Token Authentication
- ✅ Token Auto-Refresh
- ✅ Protected Routes
- ✅ Add Clothes
- ✅ Edit Clothes
- ✅ Delete Clothes
- ✅ View Clothes Grid
- ✅ Filter Clothes
- ✅ Generate Outfits
- ✅ View Outfits
- ✅ Delete Outfits
- ✅ Responsive Design
- ✅ Loading States
- ✅ Error Handling
- ✅ Form Validation
- ✅ Smooth Animations
- ✅ Glassmorphism UI
- ✅ Mobile-Friendly

## 🎨 UI/UX Features

- **Premium Design**: Glassmorphism effects, soft shadows, gradients
- **Smooth Animations**: Transitions, hover effects, loading spinners
- **Responsive**: Works on all screen sizes
- **Intuitive**: Clear navigation, obvious actions
- **Modern**: Clean, minimal, elegant
- **Accessible**: Good contrast, readable fonts
- **Fast**: Optimized performance, lazy loading

## 📞 Support

If you encounter any issues:

1. Check this guide first
2. Check browser console (F12) for errors
3. Check backend logs for API errors
4. Verify database is running
5. Ensure all migrations are applied

## 🎉 Enjoy Your Wardrobe App!

You now have a fully functional, beautiful wardrobe management application. Start adding your clothes and generating amazing outfits!

