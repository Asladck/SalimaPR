package repository

import (
	"SalimProject/models"
	"gorm.io/gorm"
)

type OutfitRepository struct {
	db *gorm.DB
}

func NewOutfitRepository(db *gorm.DB) *OutfitRepository {
	return &OutfitRepository{db: db}
}

func (r *OutfitRepository) CreateOutfit(o *models.Outfit) error {
	return r.db.Create(o).Error
}

func (r *OutfitRepository) GetUserClothesByID(userId string, category string) ([]models.Clothes, error) {
	var items []models.Clothes
	err := r.db.Where("user_id = ? AND category = ?", userId, category).Find(&items).Error
	return items, err
}

func (r *OutfitRepository) GetAllOutfits(userId string) ([]models.Outfit, error) {
	var outfits []models.Outfit
	err := r.db.Where("user_id = ?", userId).Find(&outfits).Error
	return outfits, err
}
func (r *OutfitRepository) GetOutfitById(userId, outId string) (models.Outfit, error) {
	var outfit models.Outfit
	err := r.db.Where("user_id = ? AND id = ?", userId, outId).First(&outfit).Error
	return outfit, err
}
func (r *OutfitRepository) DeleteOutfitById(userId, outId string) error {
	return r.db.Where("user_id = ? AND id = ?", userId, outId).Delete(&models.Outfit{}).Error
}
