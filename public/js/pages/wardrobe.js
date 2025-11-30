// Wardrobe Page Controller
import { protectPage, logout } from '../helpers/auth.js';
import { getAllClothes, addClothes, updateClothes, deleteClothes } from '../api/clothes.js';

// Protect this page
protectPage();

// State
let allClothes = [];
let currentEditId = null;
let deleteItemId = null;

// DOM Elements
const logoutBtn = document.getElementById('logoutBtn');
const addClothesBtn = document.getElementById('addClothesBtn');
const clothesGrid = document.getElementById('clothesGrid');
const loadingState = document.getElementById('loadingState');
const emptyState = document.getElementById('emptyState');
const categoryFilter = document.getElementById('categoryFilter');
const seasonFilter = document.getElementById('seasonFilter');

// Modal elements
const clothesModal = document.getElementById('clothesModal');
const deleteModal = document.getElementById('deleteModal');
const clothesForm = document.getElementById('clothesForm');
const modalTitle = document.getElementById('modalTitle');
const closeModal = document.getElementById('closeModal');
const cancelBtn = document.getElementById('cancelBtn');
const submitBtn = document.getElementById('submitBtn');
const modalError = document.getElementById('modalError');

// Delete modal elements
const closeDeleteModal = document.getElementById('closeDeleteModal');
const cancelDeleteBtn = document.getElementById('cancelDeleteBtn');
const confirmDeleteBtn = document.getElementById('confirmDeleteBtn');

// Initialize
document.addEventListener('DOMContentLoaded', () => {
    loadClothes();
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

    // Add clothes button
    addClothesBtn?.addEventListener('click', () => {
        openAddModal();
    });

    // Modal close buttons
    closeModal?.addEventListener('click', closeClothesModal);
    cancelBtn?.addEventListener('click', closeClothesModal);
    closeDeleteModal?.addEventListener('click', closeDeleteConfirmModal);
    cancelDeleteBtn?.addEventListener('click', closeDeleteConfirmModal);

    // Click outside modal to close
    clothesModal?.addEventListener('click', (e) => {
        if (e.target === clothesModal) {
            closeClothesModal();
        }
    });

    deleteModal?.addEventListener('click', (e) => {
        if (e.target === deleteModal) {
            closeDeleteConfirmModal();
        }
    });

    // Form submission
    clothesForm?.addEventListener('submit', handleFormSubmit);

    // Delete confirmation
    confirmDeleteBtn?.addEventListener('click', handleDelete);

    // Filters
    categoryFilter?.addEventListener('change', filterClothes);
    seasonFilter?.addEventListener('change', filterClothes);
}

/**
 * Load all clothes
 */
async function loadClothes() {
    showLoading();

    try {
        const response = await getAllClothes();
        allClothes = response.data || response || [];
        filterClothes();
    } catch (error) {
        console.error('Failed to load clothes:', error);
        showEmpty();
    }
}

/**
 * Filter and render clothes
 */
function filterClothes() {
    const category = categoryFilter.value;
    const season = seasonFilter.value;

    let filtered = allClothes;

    if (category) {
        filtered = filtered.filter(item => item.category === category);
    }

    if (season) {
        filtered = filtered.filter(item => item.season === season || item.season === 'all');
    }

    renderClothes(filtered);
}

/**
 * Render clothes grid
 */
function renderClothes(clothes) {
    hideLoading();

    if (!clothes || clothes.length === 0) {
        showEmpty();
        return;
    }

    hideEmpty();

    clothesGrid.innerHTML = clothes.map(item => `
        <div class="clothes-card" data-id="${item.id}">
            <div class="clothes-image-wrapper">
                ${item.image_url 
                    ? `<img src="${item.image_url}" alt="${item.name}" class="clothes-image" onerror="this.style.display='none'">` 
                    : `<div class="clothes-image-placeholder">👕</div>`
                }
            </div>
            <div class="clothes-info">
                <h3 class="clothes-name">${escapeHtml(item.name)}</h3>
                <div class="clothes-details">
                    <span class="clothes-tag clothes-tag--category">${escapeHtml(item.category)}</span>
                    <span class="clothes-tag clothes-tag--season">${escapeHtml(item.season)}</span>
                    <span class="clothes-tag">${escapeHtml(item.color)}</span>
                </div>
                <div class="clothes-actions">
                    <button class="clothes-action-btn clothes-action-btn--edit" onclick="window.editClothes(${item.id})">
                        Edit
                    </button>
                    <button class="clothes-action-btn clothes-action-btn--delete" onclick="window.confirmDelete(${item.id})">
                        Delete
                    </button>
                </div>
            </div>
        </div>
    `).join('');
}

/**
 * Open add modal
 */
function openAddModal() {
    currentEditId = null;
    modalTitle.textContent = 'Add Clothing Item';
    clothesForm.reset();
    document.getElementById('clothesId').value = '';
    hideModalError();
    clothesModal.classList.add('active');
}

/**
 * Open edit modal
 */
window.editClothes = function(id) {
    const item = allClothes.find(c => c.id === id);
    if (!item) return;

    currentEditId = id;
    modalTitle.textContent = 'Edit Clothing Item';

    document.getElementById('clothesId').value = item.id;
    document.getElementById('clothesName').value = item.name;
    document.getElementById('category').value = item.category;
    document.getElementById('color').value = item.color;
    document.getElementById('season').value = item.season;
    document.getElementById('material').value = item.material;
    document.getElementById('imageUrl').value = item.image_url || '';

    hideModalError();
    clothesModal.classList.add('active');
};

/**
 * Close clothes modal
 */
function closeClothesModal() {
    clothesModal.classList.remove('active');
    clothesForm.reset();
    currentEditId = null;
}

/**
 * Handle form submission
 */
async function handleFormSubmit(e) {
    e.preventDefault();
    hideModalError();

    const formData = {
        name: document.getElementById('clothesName').value.trim(),
        category: document.getElementById('category').value,
        color: document.getElementById('color').value.trim(),
        season: document.getElementById('season').value,
        material: document.getElementById('material').value.trim(),
        image_url: document.getElementById('imageUrl').value.trim(),
    };

    setFormLoading(true);

    try {
        if (currentEditId) {
            await updateClothes(currentEditId, formData);
        } else {
            await addClothes(formData);
        }

        closeClothesModal();
        await loadClothes();
    } catch (error) {
        showModalError(error.message || 'Failed to save item');
        setFormLoading(false);
    }
}

/**
 * Confirm delete
 */
window.confirmDelete = function(id) {
    deleteItemId = id;
    deleteModal.classList.add('active');
};

/**
 * Close delete modal
 */
function closeDeleteConfirmModal() {
    deleteModal.classList.remove('active');
    deleteItemId = null;
}

/**
 * Handle delete
 */
async function handleDelete() {
    if (!deleteItemId) return;

    setDeleteLoading(true);

    try {
        await deleteClothes(deleteItemId);
        closeDeleteConfirmModal();
        await loadClothes();
    } catch (error) {
        alert(error.message || 'Failed to delete item');
        setDeleteLoading(false);
    }
}

/**
 * Show/hide loading state
 */
function showLoading() {
    loadingState.style.display = 'flex';
    clothesGrid.style.display = 'none';
    emptyState.style.display = 'none';
}

function hideLoading() {
    loadingState.style.display = 'none';
    clothesGrid.style.display = 'grid';
}

/**
 * Show/hide empty state
 */
function showEmpty() {
    emptyState.style.display = 'flex';
    clothesGrid.style.display = 'none';
}

function hideEmpty() {
    emptyState.style.display = 'none';
}

/**
 * Show/hide modal error
 */
function showModalError(message) {
    modalError.textContent = message;
    modalError.style.display = 'block';
}

function hideModalError() {
    modalError.style.display = 'none';
    modalError.textContent = '';
}

/**
 * Set form loading state
 */
function setFormLoading(loading) {
    const btnText = submitBtn.querySelector('.btn-text');
    const btnLoader = submitBtn.querySelector('.btn-loader');

    if (loading) {
        submitBtn.disabled = true;
        btnText.style.display = 'none';
        btnLoader.style.display = 'block';
    } else {
        submitBtn.disabled = false;
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

