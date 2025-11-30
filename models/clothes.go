package models

import "time"

type Clothes struct {
	Id       string `db:"id"`
	UserId   string `db:"user_id"`
	Name     string `db:"name"`
	Category string `db:"category"`
	Color    string `db:"color"`
	Season   string `db:"season"`
	Material string `db:"material"`
	ImageURL string `db:"image_url"`

	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
