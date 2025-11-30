# Wardrobe Outfit Generator - Frontend

A beautiful, modern frontend application for managing your wardrobe and generating outfit combinations.

## 🎨 Features

- **User Authentication**: Secure sign up, login, and token-based authentication
- **Wardrobe Management**: Add, edit, delete, and organize your clothing items
- **Outfit Generation**: AI-powered outfit suggestions based on your wardrobe
- **Responsive Design**: Works seamlessly on desktop, tablet, and mobile devices
- **Premium UI**: Glassmorphism effects, smooth animations, and elegant design

## 📁 Project Structure

```
public/
├── index.html              # Home page
├── login.html              # Login page
├── signup.html             # Registration page
├── wardrobe.html           # Wardrobe management page
├── outfit.html             # Outfit generator page
├── css/
│   ├── styles.css          # Global styles and theme
│   └── pages/
│       ├── auth.css        # Authentication pages styles
│       ├── home.css        # Home page styles
│       ├── wardrobe.css    # Wardrobe page styles
│       └── outfit.css      # Outfit page styles
└── js/
    ├── config.js           # API configuration
    ├── api/
    │   ├── auth.js         # Authentication API calls
    │   ├── clothes.js      # Clothes API calls
    │   └── outfit.js       # Outfit API calls
    ├── helpers/
    │   └── auth.js         # Authentication helper functions
    └── pages/
        ├── login.js        # Login page controller
        ├── signup.js       # Signup page controller
        ├── home.js         # Home page controller
        ├── wardrobe.js     # Wardrobe page controller
        └── outfit.js       # Outfit page controller
```

## 🚀 Getting Started

### Prerequisites

- A modern web browser (Chrome, Firefox, Safari, Edge)
- Backend server running (see backend documentation)

### Configuration

1. Update the API base URL in `public/js/config.js`:

```javascript
const API_CONFIG = {
    baseURL: 'http://localhost:8080', // Change to your backend URL
    // ...
};
```

### Running the Application

#### Option 1: Using Go Backend Server

If your Go backend serves static files from the `public` folder, simply start your backend server and navigate to:

```
http://localhost:8080
```

#### Option 2: Using a Local Web Server

If you want to run the frontend separately:

**Using Python:**
```bash
cd public
python -m http.server 8000
```

**Using Node.js (http-server):**
```bash
npm install -g http-server
cd public
http-server -p 8000
```

Then navigate to: `http://localhost:8000`

## 🎯 Usage Guide

### 1. Registration & Login

1. Navigate to the signup page
2. Enter your name, username, and password (min 6 characters)
3. Click "Create Account"
4. You'll be redirected to login page
5. Enter your credentials and login

### 2. Managing Your Wardrobe

**Adding Clothes:**
1. Go to "My Closet" page
2. Click "Add Item" button
3. Fill in the details:
   - Name (e.g., "Blue Denim Jacket")
   - Category (Top, Bottom, Shoes, Accessories, Outerwear)
   - Color
   - Season
   - Material
   - Image URL
4. Click "Save Item"

**Editing Clothes:**
1. Find the item in your wardrobe
2. Click "Edit" button
3. Modify the details
4. Click "Save Item"

**Deleting Clothes:**
1. Find the item in your wardrobe
2. Click "Delete" button
3. Confirm deletion

**Filtering:**
- Use the category and season filters to find specific items

### 3. Generating Outfits

1. Go to "Outfits" page
2. Optionally specify:
   - Occasion (e.g., "Business", "Casual", "Party")
   - Season preference
3. Click "Generate Outfit"
4. View your generated outfit combinations

**Viewing Outfit Details:**
- Click "View Details" on any outfit card
- See all items in the outfit
- Delete the outfit if desired

## 🎨 Design System

### Color Palette

- **Primary**: Purple (`#8B7FFF`)
- **Secondary**: Pink (`#FF9ECD`)
- **Accent**: Yellow (`#FFD17A`)
- **Success**: Green (`#7FD1AE`)
- **Danger**: Red (`#FF8B94`)

### Typography

- **Font Family**: Segoe UI, system fonts
- **Sizes**: From 0.75rem to 3rem
- **Weights**: 400 (normal), 600 (semi-bold), 700 (bold)

### Spacing

- Uses a consistent 8px grid system
- Spacing scale: xs (4px) to 3xl (64px)

### Components

- **Buttons**: Primary, Secondary, Outline, Danger variants
- **Forms**: Styled inputs, selects, and validation
- **Cards**: Glassmorphism effect with shadows
- **Modals**: Smooth animations and backdrop blur
- **Loading States**: Elegant spinners and transitions

## 🔒 Security Features

- JWT token storage in localStorage
- Automatic token refresh on 401 responses
- Protected routes (redirect to login if not authenticated)
- XSS prevention with HTML escaping
- CSRF protection (handled by backend)

## 📱 Responsive Design

The application is fully responsive with breakpoints at:
- Desktop: > 768px
- Tablet: 481px - 768px
- Mobile: ≤ 480px

## 🌐 Browser Support

- Chrome (latest)
- Firefox (latest)
- Safari (latest)
- Edge (latest)

## 🛠️ Development

### Code Style

- **CSS**: BEM naming convention
- **JavaScript**: ES6+ modules, async/await
- **HTML**: Semantic HTML5

### Best Practices

- Separation of concerns (logic, UI, API)
- Reusable components
- Error handling
- Loading states
- Form validation
- No inline styles or scripts

## 🐛 Troubleshooting

### Login/Signup Issues

**Problem**: "Failed to fetch" error
- **Solution**: Check if backend server is running
- **Solution**: Verify API base URL in `config.js`
- **Solution**: Check CORS settings on backend

### Image Upload Issues

**Problem**: Images not displaying
- **Solution**: Use direct image URLs (not file uploads)
- **Solution**: Ensure image URLs are accessible
- **Solution**: Check browser console for errors

### Token Expiration

**Problem**: Automatically logged out
- **Solution**: This is normal when access token expires
- **Solution**: Refresh token will automatically renew session
- **Solution**: If refresh token expires, you need to login again

### Outfit Generation Issues

**Problem**: "Failed to generate outfit"
- **Solution**: Add more clothes to your wardrobe (need variety)
- **Solution**: Ensure you have clothes in different categories
- **Solution**: Check backend logs for errors

## 📄 License

This project is part of the Wardrobe application.

## 👨‍💻 Author

Created for SalimProject

## 🙏 Acknowledgments

- Icons: Emoji (Unicode)
- Design inspiration: Modern fashion apps
- CSS techniques: Glassmorphism, Neumorphism

