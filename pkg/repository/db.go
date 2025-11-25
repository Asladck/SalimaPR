package repository

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewDataBase(config Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Almaty",
			config.Host, config.Username, config.Password, config.DBName, config.Port, config.SSLMode),
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	fmt.Printf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Almaty",
		config.Host, config.Username, config.Password, config.DBName, config.Port, config.SSLMode)
	return db, err
}
