package models

import "time"

type Outfit struct {
	Id       string `json:"id" db:"id" gorm:"type:uuid;primaryKey"`
	UserId   string `json:"user_id" db:"user_id"`
	TopId    string `json:"top_id" db:"top_id"`
	BottomId string `json:"bottom_id" db:"bottom_id"`
	ShoesId  string `json:"shoes_id" db:"shoes_id"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`

	// Related items (not stored in DB, populated when fetching)
	Items []Clothes `json:"items,omitempty" gorm:"-"`
}
