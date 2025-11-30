package models

import "time"

type User struct {
	Id           string    `db:"id" gorm:"type:uuid;primaryKey"`
	Email        string    `db:"email"`
	Name         string    `db:"name"`
	Username     string    `db:"username"`
	PasswordHash string    `db:"password_hash"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
