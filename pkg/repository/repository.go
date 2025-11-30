package repository

import (
	"SalimProject/models"
	"gorm.io/gorm"
)

type Repository struct {
	Auth
	Clothes
	Outfit
}
type Outfit interface {
	GenerateOutfit(userId string)
}

type Auth interface {
	CreateUser(user *models.User) (string, error)
	GetUser(username, email string) (models.User, error)
}
type Clothes interface {
	AddClothes(item models.Clothes) (string, error)
	GetClothesByUserId(userId string) ([]models.Clothes, error)
	DeleteClothes(id string) error
	UpdateClothes(item models.Clothes) (models.Clothes, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth:    NewAuthRepository(db),
		Clothes: NewClothesRepository(db),
	}
}
