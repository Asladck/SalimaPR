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
