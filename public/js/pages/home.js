// Home Page Controller
import { protectPage, logout } from '../helpers/auth.js';

// Protect this page
protectPage();

// DOM Elements
const logoutBtn = document.getElementById('logoutBtn');

// Logout handler
if (logoutBtn) {
    logoutBtn.addEventListener('click', () => {
        if (confirm('Are you sure you want to logout?')) {
            logout();
        }
    });
}

// Add smooth scroll behavior for links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            target.scrollIntoView({
                behavior: 'smooth'
            });
        }
    });
});

