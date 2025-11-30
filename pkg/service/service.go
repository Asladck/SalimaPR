package service

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"SalimProject/pkg/repository"
)

type Service struct {
	Auth
	Clothes
}
type Clothes interface {
	AddClothes(item models.Clothes) (string, error)
	GetClothesByUserId(userId string) ([]models.Clothes, error)
	DeleteClothes(id string) error
	UpdateClothes(item models.Clothes) (models.Clothes, error)
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
	}
}
