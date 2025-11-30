package service

import (
	"SalimProject/models"
	"SalimProject/pkg/dto"
	"SalimProject/pkg/repository"
	"github.com/google/uuid"
)

type ClothesService struct {
	repo repository.Clothes
}

func NewClothesService(repo repository.Clothes) *ClothesService {
	return &ClothesService{repo: repo}
}

func (s *ClothesService) AddClothes(item *models.Clothes) (string, error) {
	item.Id = uuid.New().String()
	return s.repo.AddClothes(item)
}

func (s *ClothesService) GetClothesByUserId(userID string) ([]models.Clothes, error) {
	return s.repo.GetClothesByUserId(userID)
}

func (s *ClothesService) GetClothesById(userID, itemID string) (*models.Clothes, error) {
	return s.repo.GetClothesById(userID, itemID)
}

func (s *ClothesService) DeleteClothesById(userID, itemID string) error {
	return s.repo.DeleteClothes(userID, itemID)
}
func (s *ClothesService) UpdateClothesById(userID, itemID string, input dto.ClothesUpdateInput) error {
	return s.repo.UpdateClothes(userID, itemID, input)
}
