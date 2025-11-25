package repository

import (
	"SalimProject/models"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user models.User) (string, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return "", err
	}
	return user.Id, nil
}

func (r *AuthRepository) GetUser(username, password, email string) (models.User, error) {
	var user models.User
	query := r.db.Model(&models.User{})

	if username != "" {
		query = query.Where("username = ?", username)
	}
	if password != "" {
		query = query.Where("password = ?", password)
	}
	if email != "" {
		query = query.Where("email = ?", email)
	}

	if err := query.First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
