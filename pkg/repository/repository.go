package repository

import (
	"SalimProject/models"
	"gorm.io/gorm"
)

type Repository struct {
	Auth
}
type Auth interface {
	CreateUser(user models.User) (string, error)
	GetUser(username, password, email string) (models.User, error)
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}
