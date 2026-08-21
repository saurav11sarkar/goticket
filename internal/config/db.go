package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *Config) *gorm.DB {
	loadConfig := cfg.Dsn

	db, err := gorm.Open(postgres.Open(loadConfig), &gorm.Config{
		TranslateError: true,
	})
	if err != nil {
		log.Fatal("failed to connect database", "error", err)

	}
	fmt.Printf("Connected to database")
	return db
}
