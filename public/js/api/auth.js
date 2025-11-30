// Auth API Functions
import API_CONFIG from '../config.js';
import { setTokens, getRefreshToken } from '../helpers/auth.js';

/**
 * Sign up new user
 */
export async function signUp(name, username, email, password) {
    const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.signUp}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ name, username, email, password }),
    });

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Registration failed' }));
        throw new Error(error.message || 'Registration failed');
    }

    return response.json();
}

/**
 * Sign in user
 */
export async function signIn(username, password) {
    const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.signIn}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    });

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Login failed' }));
        throw new Error(error.message || 'Invalid credentials');
    }

    const data = await response.json();

    // Store tokens (backend returns access_token and refresh_token with underscores)
    if (data.access_token && data.refresh_token) {
        setTokens(data.access_token, data.refresh_token);
    }

    return data;
}

/**
 * Refresh access token
 */
export async function refreshAccessToken() {
    const refreshToken = getRefreshToken();

    if (!refreshToken) {
        throw new Error('No refresh token available');
    }

    const response = await fetch(`${API_CONFIG.baseURL}${API_CONFIG.endpoints.refresh}`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ refresh_token: refreshToken }),
    });

    if (!response.ok) {
        throw new Error('Token refresh failed');
    }

    const data = await response.json();

    // Update tokens (backend returns access_token with underscore)
    if (data.access_token) {
        setTokens(data.access_token, data.refresh_token || refreshToken);
    }

    return data;
}

