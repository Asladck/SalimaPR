package service

import (
	"SalimProject/models"
	"SalimProject/pkg/repository"
)

type ClothesService struct {
	repo repository.ClothesRepository
}

func NewClothesService(repo repository.ClothesRepository) *ClothesService {
	return &ClothesService{repo: repo}
}

func (s *ClothesService) AddClothes(item models.Clothes) (string, error) {
	return s.repo.AddClothes(item)
}

func (s *ClothesService) GetClothesByUserId(userID string) ([]models.Clothes, error) {
	return s.repo.GetClothesByUserID(userID)
}
