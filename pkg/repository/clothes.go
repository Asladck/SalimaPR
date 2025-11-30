package repository

import (
	"SalimProject/models"
	"gorm.io/gorm"
)

type ClothesRepository struct {
	db *gorm.DB
}

func NewClothesRepository(db *gorm.DB) *ClothesRepository {
	return &ClothesRepository{db: db}
}

func (r *ClothesRepository) AddClothes(item models.Clothes) (string, error) {
	if err := r.db.Create(item).Error; err != nil {
		return "", err
	}
	return item.Id, nil
}

func (r *ClothesRepository) GetClothesByUserID(userId string) ([]models.Clothes, error) {
	var clothes []models.Clothes
	err := r.db.Where("user_id = ?", userId).Find(&clothes).Error
	return clothes, err
}
