package repository

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"gorm.io/gorm"
)

type ClothesRepository struct {
	db *gorm.DB
}

func NewClothesRepository(db *gorm.DB) *ClothesRepository {
	return &ClothesRepository{db: db}
}

func (r *ClothesRepository) AddClothes(item *models.Clothes) (string, error) {
	if err := r.db.Create(item).Error; err != nil {
		return "", err
	}
	return item.Id, nil
}

func (r *ClothesRepository) GetClothesByUserId(userId string) ([]models.Clothes, error) {
	var clothes []models.Clothes
	err := r.db.Where("user_id = ?", userId).Find(&clothes).Error
	return clothes, err
}

func (r *ClothesRepository) GetClothesById(userId, itemId string) (*models.Clothes, error) {
	var item models.Clothes
	err := r.db.Where("user_id = ? AND id = ?", userId, itemId).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *ClothesRepository) DeleteClothes(userID, itemID string) error {
	return r.db.Where("user_id = ? AND id = ?", userID, itemID).Delete(&models.Clothes{}).Error
}
func (r *ClothesRepository) UpdateClothes(userID, itemID string, input dto.ClothesUpdateInput) error {
	return r.db.Model(&models.Clothes{}).Where("user_id = ? AND id = ?", userID, itemID).Updates(input).Error
}

func (r *ClothesRepository) GetClothesByCategory(userId string, category string) ([]models.Clothes, error) {
	var clothes []models.Clothes
	err := r.db.Where("user_id = ? AND category = ?", userId, category).Find(&clothes).Error
	return clothes, err
}
