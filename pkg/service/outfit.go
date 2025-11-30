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
	outfitRepo  repository.Outfit
	clothesRepo repository.Clothes
}

func NewOutfitService(outfitRepo repository.Outfit, clothesRepo repository.Clothes) *OutfitService {
	return &OutfitService{
		outfitRepo:  outfitRepo,
		clothesRepo: clothesRepo,
	}
}

func (s *OutfitService) GenerateOutfit(userId string) (string, error) {
	tops, err := s.clothesRepo.GetClothesByCategory(userId, "top")
	if err != nil {
		return "", err
	}

	bottoms, err := s.clothesRepo.GetClothesByCategory(userId, "bottom")
	if err != nil {
		return "", err
	}

	shoes, err := s.clothesRepo.GetClothesByCategory(userId, "shoes")
	if err != nil {
		return "", err
	}

	if len(tops) == 0 || len(bottoms) == 0 || len(shoes) == 0 {
		return "", errors.New("not enough clothes to generate outfit")
	}

	rand.Seed(time.Now().UnixNano())
	top := tops[rand.Intn(len(tops))]
	bottom := bottoms[rand.Intn(len(bottoms))]
	shoe := shoes[rand.Intn(len(shoes))]

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
		return "", err
	}

	return outfit.Id, nil
}

func (s *OutfitService) GetAllOutfits(userId string) ([]models.Outfit, error) {
	outfits, err := s.outfitRepo.GetAllOutfits(userId)
	if err != nil {
		return nil, err
	}

	// Загрузить полную информацию о вещах для каждого outfit
	for i := range outfits {
		items := []models.Clothes{}

		// Загрузить top
		if outfits[i].TopId != "" {
			top, err := s.clothesRepo.GetClothesById(userId, outfits[i].TopId)
			if err == nil && top != nil {
				items = append(items, *top)
			}
		}

		// Загрузить bottom
		if outfits[i].BottomId != "" {
			bottom, err := s.clothesRepo.GetClothesById(userId, outfits[i].BottomId)
			if err == nil && bottom != nil {
				items = append(items, *bottom)
			}
		}

		// Загрузить shoes
		if outfits[i].ShoesId != "" {
			shoes, err := s.clothesRepo.GetClothesById(userId, outfits[i].ShoesId)
			if err == nil && shoes != nil {
				items = append(items, *shoes)
			}
		}

		outfits[i].Items = items
	}

	return outfits, nil
}

func (s *OutfitService) GetOutfitById(userId, outId string) (models.Outfit, error) {
	outfit, err := s.outfitRepo.GetOutfitById(userId, outId)
	if err != nil {
		return outfit, err
	}

	// Загрузить полную информацию о вещах
	items := []models.Clothes{}

	// Загрузить top
	if outfit.TopId != "" {
		top, err := s.clothesRepo.GetClothesById(userId, outfit.TopId)
		if err == nil && top != nil {
			items = append(items, *top)
		}
	}

	// Загрузить bottom
	if outfit.BottomId != "" {
		bottom, err := s.clothesRepo.GetClothesById(userId, outfit.BottomId)
		if err == nil && bottom != nil {
			items = append(items, *bottom)
		}
	}

	// Загрузить shoes
	if outfit.ShoesId != "" {
		shoes, err := s.clothesRepo.GetClothesById(userId, outfit.ShoesId)
		if err == nil && shoes != nil {
			items = append(items, *shoes)
		}
	}

	outfit.Items = items

	return outfit, nil
}
func (s *OutfitService) DeleteOutfitById(userId, outId string) error {
	return s.outfitRepo.DeleteOutfitById(userId, outId)
}
