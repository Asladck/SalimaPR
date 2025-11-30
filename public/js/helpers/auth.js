// Authentication Helper Functions

/**
 * Get access token from localStorage
 */
export function getAccessToken() {
    return localStorage.getItem('accessToken');
}

/**
 * Get refresh token from localStorage
 */
export function getRefreshToken() {
    return localStorage.getItem('refreshToken');
}

/**
 * Set tokens in localStorage
 */
export function setTokens(accessToken, refreshToken) {
    localStorage.setItem('accessToken', accessToken);
    if (refreshToken) {
        localStorage.setItem('refreshToken', refreshToken);
    }
}

/**
 * Clear tokens from localStorage
 */
export function clearTokens() {
    localStorage.removeItem('accessToken');
    localStorage.removeItem('refreshToken');
}

/**
 * Check if user is authenticated
 */
export function isAuthenticated() {
    return !!getAccessToken();
}

/**
 * Redirect to login page
 */
export function redirectToLogin() {
    clearTokens();
    window.location.href = '/login.html';
}

/**
 * Redirect to home page
 */
export function redirectToHome() {
    window.location.href = '/index.html';
}

/**
 * Protect page - redirect to login if not authenticated
 */
export function protectPage() {
    if (!isAuthenticated()) {
        redirectToLogin();
    }
}

/**
 * Redirect authenticated users away from auth pages
 */
export function redirectIfAuthenticated() {
    if (isAuthenticated()) {
        redirectToHome();
    }
}

/**
 * Logout user
 */
export function logout() {
    clearTokens();
    redirectToLogin();
}

