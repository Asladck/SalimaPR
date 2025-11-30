package models

import "time"

type User struct {
	Id           string    `json:"id" db:"id" gorm:"type:uuid;primaryKey"`
	Email        string    `json:"email" db:"email"`
	Name         string    `json:"name" db:"name"`
	Username     string    `json:"username" db:"username"`
	PasswordHash string    `json:"-" db:"password_hash"`
	CreatedAt    time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at,omitempty" db:"updated_at"`
}
