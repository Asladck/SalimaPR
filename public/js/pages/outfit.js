// Outfit Page Controller
import { protectPage, logout } from '../helpers/auth.js';
import { generateOutfit, getAllOutfits, getOutfitById, deleteOutfit } from '../api/outfit.js';

// Protect this page
protectPage();

// State
let allOutfits = [];
let currentOutfitId = null;
let deleteOutfitId = null;
let currentTab = 'all';

// DOM Elements
const logoutBtn = document.getElementById('logoutBtn');
const generateBtn = document.getElementById('generateBtn');
const generatorForm = document.getElementById('generatorForm');
const submitGenerateBtn = document.getElementById('submitGenerateBtn');
const generateError = document.getElementById('generateError');
const outfitsGrid = document.getElementById('outfitsGrid');
const loadingState = document.getElementById('loadingState');
const emptyState = document.getElementById('emptyState');

// Modal elements
const outfitModal = document.getElementById('outfitModal');
const deleteModal = document.getElementById('deleteModal');
const closeModal = document.getElementById('closeModal');
const closeDetailsBtn = document.getElementById('closeDetailsBtn');
const outfitDetails = document.getElementById('outfitDetails');
const deleteOutfitBtn = document.getElementById('deleteOutfitBtn');

// Delete modal elements
const closeDeleteModal = document.getElementById('closeDeleteModal');
const cancelDeleteBtn = document.getElementById('cancelDeleteBtn');
const confirmDeleteBtn = document.getElementById('confirmDeleteBtn');

// Tab buttons
const tabButtons = document.querySelectorAll('.tab-btn');

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    loadOutfits();
    setupEventListeners();
});

/**
 * Setup event listeners
 */
function setupEventListeners() {
    // Logout
    logoutBtn?.addEventListener('click', () => {
        if (confirm('Are you sure you want to logout?')) {
            logout();
        }
    });

    // Generate button
    generateBtn?.addEventListener('click', () => {
        document.querySelector('.generator-section').scrollIntoView({ behavior: 'smooth' });
    });

    // Generator form
    generatorForm?.addEventListener('submit', handleGenerateOutfit);

    // Tabs
    tabButtons.forEach(btn => {
        btn.addEventListener('click', () => {
            tabButtons.forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            currentTab = btn.dataset.tab;
            filterOutfits();
        });
    });

    // Modal close buttons
    closeModal?.addEventListener('click', closeOutfitModal);
    closeDetailsBtn?.addEventListener('click', closeOutfitModal);
    closeDeleteModal?.addEventListener('click', closeDeleteConfirmModal);
    cancelDeleteBtn?.addEventListener('click', closeDeleteConfirmModal);

    // Delete outfit button in detail modal
    deleteOutfitBtn?.addEventListener('click', () => {
        closeOutfitModal();
        confirmDeleteOutfit(currentOutfitId);
    });

    // Delete confirmation
    confirmDeleteBtn?.addEventListener('click', handleDeleteOutfit);

    // Click outside modal to close
    outfitModal?.addEventListener('click', (e) => {
        if (e.target === outfitModal) {
            closeOutfitModal();
        }
    });

    deleteModal?.addEventListener('click', (e) => {
        if (e.target === deleteModal) {
            closeDeleteConfirmModal();
        }
    });
}

/**
 * Load all outfits
 */
async function loadOutfits() {
    showLoading();

    try {
        const response = await getAllOutfits();
        allOutfits = response.data || response || [];
        filterOutfits();
    } catch (error) {
        console.error('Failed to load outfits:', error);
        showEmpty();
    }
}

/**
 * Filter outfits based on current tab
 */
function filterOutfits() {
    let filtered = allOutfits;

    if (currentTab === 'recent') {
        // Sort by created date and take last 10
        filtered = [...allOutfits]
            .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
            .slice(0, 10);
    }

    renderOutfits(filtered);
}

/**
 * Render outfits grid
 */
function renderOutfits(outfits) {
    hideLoading();

    if (!outfits || outfits.length === 0) {
        showEmpty();
        return;
    }

    hideEmpty();

    outfitsGrid.innerHTML = outfits.map(outfit => {
        const items = outfit.items || [];
        const date = outfit.created_at ? new Date(outfit.created_at).toLocaleDateString() : 'N/A';

        return `
            <div class="outfit-card" data-id="${outfit.id}">
                <div class="outfit-header">
                    <div class="outfit-title">Outfit #${outfit.id}</div>
                    <div class="outfit-date">${date}</div>
                </div>
                <div class="outfit-items">
                    ${items.slice(0, 4).map(item => `
                        <div class="outfit-item">
                            ${item.image_url 
                                ? `<img src="${item.image_url}" alt="${item.name}" class="outfit-item-image" onerror="this.style.display='none'">` 
                                : `<div class="outfit-item-placeholder">👕</div>`
                            }
                            <div class="outfit-item-label">${escapeHtml(item.category || 'Item')}</div>
                        </div>
                    `).join('')}
                </div>
                <div class="outfit-footer">
                    <div class="outfit-meta">
                        ${outfit.occasion ? `<span class="outfit-tag">${escapeHtml(outfit.occasion)}</span>` : ''}
                        <span class="outfit-tag">${items.length} items</span>
                    </div>
                    <div class="outfit-actions">
                        <button class="outfit-action-btn outfit-action-btn--view" onclick="window.viewOutfit(${outfit.id})">
                            View Details
                        </button>
                        <button class="outfit-action-btn outfit-action-btn--delete" onclick="window.confirmDeleteOutfit(${outfit.id})">
                            Delete
                        </button>
                    </div>
                </div>
            </div>
        `;
    }).join('');
}

/**
 * Handle outfit generation
 */
async function handleGenerateOutfit(e) {
    e.preventDefault();
    hideGenerateError();

    const preferences = {
        occasion: document.getElementById('occasion').value.trim() || undefined,
        season: document.getElementById('season').value || undefined,
    };

    // Remove undefined values
    Object.keys(preferences).forEach(key =>
        preferences[key] === undefined && delete preferences[key]
    );

    setGenerateLoading(true);

    try {
        await generateOutfit(preferences);
        generatorForm.reset();
        await loadOutfits();

        // Scroll to outfits
        document.querySelector('.outfits-grid').scrollIntoView({ behavior: 'smooth' });
    } catch (error) {
        showGenerateError(error.message || 'Failed to generate outfit. Make sure you have enough clothes in your wardrobe.');
    } finally {
        setGenerateLoading(false);
    }
}

/**
 * View outfit details
 */
window.viewOutfit = async function(id) {
    currentOutfitId = id;

    try {
        const outfit = await getOutfitById(id);
        displayOutfitDetails(outfit);
        outfitModal.classList.add('active');
    } catch (error) {
        alert(error.message || 'Failed to load outfit details');
    }
};

/**
 * Display outfit details in modal
 */
function displayOutfitDetails(outfit) {
    const items = outfit.items || outfit.data?.items || [];
    const date = outfit.created_at ? new Date(outfit.created_at).toLocaleDateString() : 'N/A';

    outfitDetails.innerHTML = `
        <div class="outfit-details-header">
            <h3 class="outfit-details-title">Outfit #${outfit.id}</h3>
            <p class="outfit-details-info">Created on ${date}</p>
            ${outfit.occasion ? `<p class="outfit-details-info">Occasion: ${escapeHtml(outfit.occasion)}</p>` : ''}
        </div>
        <div class="outfit-details-items">
            ${items.map(item => `
                <div class="outfit-detail-item">
                    <div class="outfit-detail-image-wrapper">
                        ${item.image_url 
                            ? `<img src="${item.image_url}" alt="${item.name}" class="outfit-detail-image" onerror="this.style.display='none'">` 
                            : `<div class="clothes-image-placeholder">👕</div>`
                        }
                    </div>
                    <div class="outfit-detail-name">${escapeHtml(item.name)}</div>
                    <div class="outfit-detail-category">${escapeHtml(item.category)}</div>
                </div>
            `).join('')}
        </div>
    `;
}

/**
 * Close outfit detail modal
 */
function closeOutfitModal() {
    outfitModal.classList.remove('active');
    currentOutfitId = null;
}

/**
 * Confirm delete outfit
 */
window.confirmDeleteOutfit = function(id) {
    deleteOutfitId = id;
    deleteModal.classList.add('active');
};

/**
 * Close delete confirmation modal
 */
function closeDeleteConfirmModal() {
    deleteModal.classList.remove('active');
    deleteOutfitId = null;
}

/**
 * Handle delete outfit
 */
async function handleDeleteOutfit() {
    if (!deleteOutfitId) return;

    setDeleteLoading(true);

    try {
        await deleteOutfit(deleteOutfitId);
        closeDeleteConfirmModal();
        await loadOutfits();
    } catch (error) {
        alert(error.message || 'Failed to delete outfit');
    } finally {
        setDeleteLoading(false);
    }
}

/**
 * Show/hide loading state
 */
function showLoading() {
    loadingState.style.display = 'flex';
    outfitsGrid.style.display = 'none';
    emptyState.style.display = 'none';
}

function hideLoading() {
    loadingState.style.display = 'none';
    outfitsGrid.style.display = 'grid';
}

/**
 * Show/hide empty state
 */
function showEmpty() {
    emptyState.style.display = 'flex';
    outfitsGrid.style.display = 'none';
}

function hideEmpty() {
    emptyState.style.display = 'none';
}

/**
 * Show/hide generate error
 */
function showGenerateError(message) {
    generateError.textContent = message;
    generateError.style.display = 'block';
}

function hideGenerateError() {
    generateError.style.display = 'none';
    generateError.textContent = '';
}

/**
 * Set generate loading state
 */
function setGenerateLoading(loading) {
    const btnText = submitGenerateBtn.querySelector('.btn-text');
    const btnLoader = submitGenerateBtn.querySelector('.btn-loader');

    if (loading) {
        submitGenerateBtn.disabled = true;
        btnText.style.display = 'none';
        btnLoader.style.display = 'block';
    } else {
        submitGenerateBtn.disabled = false;
        btnText.style.display = 'block';
        btnLoader.style.display = 'none';
    }
}

/**
 * Set delete loading state
 */
function setDeleteLoading(loading) {
    const btnText = confirmDeleteBtn.querySelector('.btn-text');
    const btnLoader = confirmDeleteBtn.querySelector('.btn-loader');

    if (loading) {
        confirmDeleteBtn.disabled = true;
        btnText.style.display = 'none';
        btnLoader.style.display = 'block';
    } else {
        confirmDeleteBtn.disabled = false;
        btnText.style.display = 'block';
        btnLoader.style.display = 'none';
    }
}

/**
 * Escape HTML to prevent XSS
 */
function escapeHtml(text) {
    const div = document.createElement('div');
    div.textContent = text;
    return div.innerHTML;
}

