// Sign Up Page Controller
import { signUp } from '../api/auth.js';
import { redirectIfAuthenticated } from '../helpers/auth.js';

// Check if already authenticated
redirectIfAuthenticated();

// DOM Elements
const signupForm = document.getElementById('signupForm');
const nameInput = document.getElementById('name');
const usernameInput = document.getElementById('username');
const passwordInput = document.getElementById('password');
const signupBtn = document.getElementById('signupBtn');
const errorMessage = document.getElementById('errorMessage');

// Form submission handler
signupForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    // Clear previous errors
    hideError();

    // Get form values
    const name = nameInput.value.trim();
    const username = usernameInput.value.trim();
    const email = document.getElementById('email').value.trim();
    const password = passwordInput.value;

    // Validate
    if (!name || !username || !email || !password) {
        showError('Please fill in all fields');
        return;
    }

    // Validate email format
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
        showError('Please enter a valid email address');
        return;
    }

    if (password.length < 6) {
        showError('Password must be at least 6 characters long');
        return;
    }

    if (username.length < 3) {
        showError('Username must be at least 3 characters long');
        return;
    }

    // Show loading state
    setLoading(true);

    try {
        // Attempt signup
        await signUp(name, username, email, password);

        // Show success and redirect
        alert('Account created successfully! Please log in.');
        window.location.href = '/login.html';
    } catch (error) {
        showError(error.message || 'Registration failed. Please try again.');
        setLoading(false);
    }
});

/**
 * Show error message
 */
function showError(message) {
    errorMessage.textContent = message;
    errorMessage.style.display = 'block';
}

/**
 * Hide error message
 */
function hideError() {
    errorMessage.style.display = 'none';
    errorMessage.textContent = '';
}

/**
 * Set loading state
 */
function setLoading(loading) {
    const btnText = signupBtn.querySelector('.btn-text');
    const btnLoader = signupBtn.querySelector('.btn-loader');

    if (loading) {
        signupBtn.disabled = true;
        btnText.style.display = 'none';
        btnLoader.style.display = 'block';
    } else {
        signupBtn.disabled = false;
        btnText.style.display = 'block';
        btnLoader.style.display = 'none';
    }
}

