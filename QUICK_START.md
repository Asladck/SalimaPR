# 🚀 QUICK START GUIDE - Wardrobe Outfit Generator

## ⚡ 3-Minute Setup

### Step 1: Start Server (30 seconds)
```bash
cd C:\Users\DWA\GolandProjects\SalimProject
go run cmd/main.go
```
✅ Server running on http://localhost:8080

### Step 2: Open Browser (10 seconds)
```
http://localhost:8080
```
✅ You'll see the beautiful home page

### Step 3: Create Account (1 minute)
1. Click "Sign up" link
2. Enter: Name, Username, Password
3. Click "Create Account"
4. Login with your credentials
✅ You're now logged in!

### Step 4: Add Clothes (1 minute)
1. Go to "My Closet"
2. Click "Add Item"
3. Use these test images:
   ```
   Top: https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=500
   Bottom: https://images.unsplash.com/photo-1542272454315-7f6fabf03775?w=500
   Shoes: https://images.unsplash.com/photo-1549298916-b41d501d3772?w=500
   ```
4. Fill name, category, color, season, material
5. Click "Save Item"
✅ Your wardrobe is ready!

### Step 5: Generate Outfit (30 seconds)
1. Go to "Outfits" page
2. Click "Generate Outfit"
3. View your styled combination!
✅ Done! You're using the app!

---

## 📋 File Checklist

### ✅ HTML Files (5)
- `public/index.html` - Home
- `public/login.html` - Login
- `public/signup.html` - Signup
- `public/wardrobe.html` - Wardrobe
- `public/outfit.html` - Outfits

### ✅ CSS Files (5)
- `public/css/styles.css` - Global
- `public/css/pages/auth.css` - Auth pages
- `public/css/pages/home.css` - Home
- `public/css/pages/wardrobe.css` - Wardrobe
- `public/css/pages/outfit.css` - Outfits

### ✅ JavaScript Files (10)
- `public/js/config.js` - Config
- `public/js/api/auth.js` - Auth API
- `public/js/api/clothes.js` - Clothes API
- `public/js/api/outfit.js` - Outfit API
- `public/js/helpers/auth.js` - Auth helpers
- `public/js/pages/login.js` - Login logic
- `public/js/pages/signup.js` - Signup logic
- `public/js/pages/home.js` - Home logic
- `public/js/pages/wardrobe.js` - Wardrobe logic
- `public/js/pages/outfit.js` - Outfit logic

---

## 🎨 Key Features

✅ User registration & login  
✅ JWT authentication  
✅ Add/edit/delete clothes  
✅ Filter by category & season  
✅ Generate outfit combinations  
✅ View outfit details  
✅ Delete outfits  
✅ Responsive design  
✅ Beautiful UI with glassmorphism  
✅ Smooth animations  

---

## 🎯 API Endpoints

### Auth
- `POST /auth/sign-up` - Register
- `POST /auth/sign-in` - Login
- `POST /auth/refresh` - Refresh token

### Clothes (requires auth)
- `GET /api/clothes/item` - List all
- `GET /api/clothes/item/:id` - Get one
- `POST /api/clothes/item` - Add
- `PUT /api/clothes/item/:id` - Update
- `DELETE /api/clothes/item/:id` - Delete

### Outfits (requires auth)
- `POST /api/outfit/generate` - Generate
- `GET /api/outfit/` - List all
- `GET /api/outfit/:id` - Get one
- `DELETE /api/outfit/:id` - Delete

---

## 🐛 Troubleshooting

**"Failed to fetch"**
→ Make sure backend is running: `go run cmd/main.go`

**Images not showing**
→ Use direct image URLs (Unsplash recommended)

**Can't login**
→ Check PostgreSQL is running and configured

**Page redirects to login**
→ Your session expired, login again

---

## 📚 Documentation

📖 **Main README**: `README.md`  
🔧 **Setup Guide**: `SETUP_GUIDE.md`  
🎨 **Design Docs**: `DESIGN_PREVIEW.md`  
✅ **Summary**: `PROJECT_SUMMARY.md`  
💡 **Frontend**: `public/README.md`  

---

## 🎨 Color Reference

```css
Primary:   #8B7FFF (Purple)
Secondary: #FF9ECD (Pink)
Accent:    #FFD17A (Yellow)
Success:   #7FD1AE (Green)
Danger:    #FF8B94 (Red)
```

---

## 💡 Pro Tips

1. **Test Images**: Use Unsplash for free quality images
2. **Categories**: Add variety (tops, bottoms, shoes, accessories)
3. **Seasons**: Match clothes to seasons for better outfits
4. **Filters**: Use category/season filters to find items quickly
5. **Mobile**: Test on phone - it's fully responsive!

---

## ✅ Status

**Frontend**: ✅ 100% Complete  
**Backend Integration**: ✅ Ready  
**Documentation**: ✅ Complete  
**Testing**: ✅ Ready to test  
**Production**: ✅ Ready to deploy  

---

## 🎉 You're All Set!

Run the server, open your browser, and enjoy your beautiful wardrobe app!

**Happy styling!** 👗✨

