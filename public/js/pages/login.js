// Login Page Controller
import { signIn } from '../api/auth.js';
import { redirectIfAuthenticated, redirectToHome } from '../helpers/auth.js';

// Check if already authenticated
redirectIfAuthenticated();

// DOM Elements
const loginForm = document.getElementById('loginForm');
const usernameInput = document.getElementById('username');
const passwordInput = document.getElementById('password');
const loginBtn = document.getElementById('loginBtn');
const errorMessage = document.getElementById('errorMessage');

// Form submission handler
loginForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    // Clear previous errors
    hideError();

    // Get form values
    const username = usernameInput.value.trim();
    const password = passwordInput.value;

    // Validate
    if (!username || !password) {
        showError('Please fill in all fields');
        return;
    }

    // Show loading state
    setLoading(true);

    try {
        // Attempt login
        await signIn(username, password);

        // Redirect to home on success
        redirectToHome();
    } catch (error) {
        showError(error.message || 'Login failed. Please try again.');
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
    const btnText = loginBtn.querySelector('.btn-text');
    const btnLoader = loginBtn.querySelector('.btn-loader');

    if (loading) {
        loginBtn.disabled = true;
        btnText.style.display = 'none';
        btnLoader.style.display = 'block';
    } else {
        loginBtn.disabled = false;
        btnText.style.display = 'block';
        btnLoader.style.display = 'none';
    }
}

