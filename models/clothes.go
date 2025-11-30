package models

import "time"

type Clothes struct {
	Id        string    `json:"id" db:"id" gorm:"type:uuid;primaryKey"`
	UserId    string    `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	Category  string    `json:"category" db:"category"`
	Color     string    `json:"color" db:"color"`
	Season    string    `json:"season" db:"season"`
	Material  string    `json:"material" db:"material"`
	ImageURL  string    `json:"image_url" db:"image_url"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
