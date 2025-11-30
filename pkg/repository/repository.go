package repository

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"gorm.io/gorm"
)

type Repository struct {
	Auth
	Clothes
	Outfit
}
type Outfit interface {
	CreateOutfit(outfit *models.Outfit) error
	GetAllOutfits(userId string) ([]models.Outfit, error)
	GetOutfitById(userId, outId string) (models.Outfit, error)
	DeleteOutfitById(userId, outId string) error
}
type Auth interface {
	CreateUser(user *models.User) (string, error)
	GetUser(username, email string) (models.User, error)
}

type Clothes interface {
	AddClothes(item *models.Clothes) (string, error)
	GetClothesByUserId(userId string) ([]models.Clothes, error)
	GetClothesById(userId, itemId string) (*models.Clothes, error)
	DeleteClothes(userID, id string) error
	UpdateClothes(userID, itemID string, item dto.ClothesUpdateInput) error
	GetClothesByCategory(userId, category string) ([]models.Clothes, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth:    NewAuthRepository(db),
		Clothes: NewClothesRepository(db),
		Outfit:  NewOutfitRepository(db),
	}
}
