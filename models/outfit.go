package models

import "time"

type Outfit struct {
	Id     string `db:"id"`
	UserId string `db:"user_id"`

	TopId    string `db:"top_id"`
	BottomId string `db:"bottom_id"`
	ShoesId  string `db:"shoes_id"`

	CreatedAt time.Time `db:"created_at"`
}
