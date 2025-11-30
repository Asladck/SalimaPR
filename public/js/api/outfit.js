// Outfit API Functions
import API_CONFIG from '../config.js';
import { getAccessToken, redirectToLogin } from '../helpers/auth.js';
import { refreshAccessToken } from './auth.js';

/**
 * Make authenticated request with auto token refresh
 */
async function authenticatedFetch(url, options = {}) {
    const token = getAccessToken();

    if (!token) {
        redirectToLogin();
        throw new Error('No authentication token');
    }

    // Add authorization header
    const headers = {
        ...options.headers,
        'Authorization': `Bearer ${token}`,
    };

    let response = await fetch(url, { ...options, headers });

    // If unauthorized, try to refresh token
    if (response.status === 401) {
        try {
            await refreshAccessToken();
            const newToken = getAccessToken();
            headers['Authorization'] = `Bearer ${newToken}`;
            response = await fetch(url, { ...options, headers });
        } catch (error) {
            redirectToLogin();
            throw error;
        }
    }

    return response;
}

/**
 * Generate new outfit
 */
export async function generateOutfit(preferences = {}) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.outfitGenerate}`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(preferences),
        }
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to generate outfit' }));
        throw new Error(error.message || 'Failed to generate outfit');
    }

    return response.json();
}

/**
 * Get all outfits
 */
export async function getAllOutfits() {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.outfits}`
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to fetch outfits' }));
        throw new Error(error.message || 'Failed to fetch outfits');
    }

    return response.json();
}

/**
 * Get outfit by ID
 */
export async function getOutfitById(id) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.outfitById(id)}`
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to fetch outfit' }));
        throw new Error(error.message || 'Failed to fetch outfit');
    }

    return response.json();
}

/**
 * Delete outfit
 */
export async function deleteOutfit(id) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.outfitById(id)}`,
        {
            method: 'DELETE',
        }
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to delete outfit' }));
        throw new Error(error.message || 'Failed to delete outfit');
    }

    return response.json();
}

