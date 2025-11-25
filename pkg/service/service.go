package service

import (
	"SalimProject/models"
	"SalimProject/pkg/repository"
)

type Service struct {
	Auth
}
type Auth interface {
	CreateUser(user models.User) (string, error)
	GenerateToken(username, password, email string) (string, string, error)
	ParseRefToken(tokenR string) (string, error)
	ParseToken(token string) (string, error)
	GenerateAccToken(userId string) (string, error)
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(rep.Auth),
	}
}
