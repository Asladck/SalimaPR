package dto

import "SalimProject/models"

type SignUpInput struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email" `
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AddClothesInput struct {
	Name     string `json:"name" binding:"required"`
	Category string `json:"category" binding:"required"`
	Color    string `json:"color"`
	Season   string `json:"season"`
	Material string `json:"material"`
	ImageURL string `json:"image_url"`
}
type ClothesUpdateInput struct {
	Name     *string `json:"name"`
	Category *string `json:"category"`
	Color    *string `json:"color"`
	Season   *string `json:"season"`
	Material *string `json:"material"`
	ImageURL *string `json:"image_url"`
}
type GetAllClothResponse struct {
	Data []models.Clothes `json:"data"`
}
type GetAllOutFitsResponse struct {
	Data []models.Outfit `json:"data"`
}
type GetOutFitResponse struct {
	Data models.Outfit `json:"data"`
}
