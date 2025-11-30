// API Configuration
const API_CONFIG = {
    baseURL: 'http://localhost:8080', // Change this to your backend URL
    endpoints: {
        // Auth
        signUp: '/auth/sign-up',
        signIn: '/auth/sign-in',
        refresh: '/auth/refresh',

        // Clothes
        clothes: '/api/clothes/item',
        clothesById: (id) => `/api/clothes/item/${id}`,

        // Outfit
        outfitGenerate: '/api/outfit/generate',
        outfits: '/api/outfit/',
        outfitById: (id) => `/api/outfit/${id}`,
    }
};

export default API_CONFIG;

