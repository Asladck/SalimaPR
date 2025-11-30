package service

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"SalimProject/pkg/repository"
)

type Service struct {
	Auth
	Clothes
	Outfit
}
type Clothes interface {
	AddClothes(item *models.Clothes) (string, error)
	GetClothesByUserId(userId string) ([]models.Clothes, error)
	GetClothesById(userId, itemId string) (*models.Clothes, error)
	DeleteClothesById(userId, id string) error
	UpdateClothesById(userId, itemId string, item dto.ClothesUpdateInput) error
}
type Outfit interface {
	GenerateOutfit(userId string) (string, error)
	GetAllOutfits(userId string) ([]models.Outfit, error)
	GetOutfitById(userId, outId string) (models.Outfit, error)
	DeleteOutfitById(userId, outId string) error
}
type Auth interface {
	CreateUser(user dto.SignUpInput) (string, error)
	GenerateToken(userId string) (string, string, error)
	ParseRefresh(tokenR string) (string, error)
	ParseToken(token string) (string, error)
	GenerateAccToken(userId string) (string, error)
	SignIn(input dto.SignInInput) (string, string, error)
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Auth:    NewAuthService(rep.Auth),
		Clothes: NewClothesService(rep.Clothes),
		Outfit:  NewOutfitService(rep.Outfit, rep.Clothes),
	}
}
