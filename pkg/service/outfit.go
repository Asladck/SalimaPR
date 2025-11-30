package service

import (
	"SalimProject/models"
	"SalimProject/pkg/repository"
	"errors"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type OutfitService struct {
	outfitRepo  repository.OutfitRepository
	clothesRepo repository.ClothesRepository
}

func NewOutfitService(outfitRepo repository.OutfitRepository, clothesRepo repository.ClothesRepository) *OutfitService {
	return &OutfitService{
		outfitRepo:  outfitRepo,
		clothesRepo: clothesRepo,
	}
}
func (s *OutfitService) GenerateOutfit(userId string) (*models.Outfit, error) {
	tops, err := s.clothesRepo.GetClothesByCategory(userId, "top")
	if err != nil {
		return nil, err
	}

	bottoms, err := s.clothesRepo.GetClothesByCategory(userId, "bottom")
	if err != nil {
		return nil, err
	}

	shoes, err := s.clothesRepo.GetClothesByCategory(userId, "shoes")
	if err != nil {
		return nil, err
	}

	if len(tops) == 0 || len(bottoms) == 0 || len(shoes) == 0 {
		return nil, errors.New("not enough clothes to generate outfit")
	}

	// Выбрать случайные элементы
	rand.Seed(time.Now().UnixNano())
	top := tops[rand.Intn(len(tops))]
	bottom := bottoms[rand.Intn(len(bottoms))]
	shoe := shoes[rand.Intn(len(shoes))]

	// Сформировать outfit
	outfit := &models.Outfit{
		Id:        uuid.New().String(),
		UserId:    userId,
		TopId:     top.Id,
		BottomId:  bottom.Id,
		ShoesId:   shoe.Id,
		CreatedAt: time.Now(),
	}

	// Сохранить в БД
	if err := s.outfitRepo.CreateOutfit(outfit); err != nil {
		return nil, err
	}

	return outfit, nil
}
func (s *OutfitService) GetUserOutfits(userId string) ([]models.Outfit, error) {
	return s.outfitRepo.GetOutfitsByUser(userId)
}
