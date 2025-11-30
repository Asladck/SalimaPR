// Clothes API Functions
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
 * Get all clothes for current user
 */
export async function getAllClothes() {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.clothes}`
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to fetch clothes' }));
        throw new Error(error.message || 'Failed to fetch clothes');
    }

    return response.json();
}

/**
 * Get clothes item by ID
 */
export async function getClothesById(id) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.clothesById(id)}`
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to fetch item' }));
        throw new Error(error.message || 'Failed to fetch item');
    }

    return response.json();
}

/**
 * Add new clothes item
 */
export async function addClothes(clothesData) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.clothes}`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(clothesData),
        }
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to add item' }));
        throw new Error(error.message || 'Failed to add item');
    }

    return response.json();
}

/**
 * Update clothes item
 */
export async function updateClothes(id, clothesData) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.clothesById(id)}`,
        {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(clothesData),
        }
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to update item' }));
        throw new Error(error.message || 'Failed to update item');
    }

    return response.json();
}

/**
 * Delete clothes item
 */
export async function deleteClothes(id) {
    const response = await authenticatedFetch(
        `${API_CONFIG.baseURL}${API_CONFIG.endpoints.clothesById(id)}`,
        {
            method: 'DELETE',
        }
    );

    if (!response.ok) {
        const error = await response.json().catch(() => ({ message: 'Failed to delete item' }));
        throw new Error(error.message || 'Failed to delete item');
    }

    return response.json();
}

