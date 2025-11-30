package dto

type SignUpInput struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type SignInInput struct {
	Email    string `json:"email" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type RefreshInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type AddClothesInput struct {
	Name     string `form:"name" binding:"required"`
	Category string `form:"category" binding:"required"`
	Color    string `form:"color"`
	Season   string `form:"season"`
	Material string `form:"material"`
}
