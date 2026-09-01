package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port string

	// Database
	DSN string

	// Cloudinary
	CloudinaryCloudName string
	CloudinaryAPIKey    string
	CloudinaryAPISecret string

	// Email
	EmailExpires string
	EmailHost    string
	EmailPort    string
	EmailAddress string
	EmailPass    string
	EmailFrom    string
	AdminEmail   string
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load(".env")

	config := &Config{
		Port: getEnv("PORT", "8080"),

		DSN: os.Getenv("DSN"),

		CloudinaryCloudName: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CloudinaryAPIKey:    os.Getenv("CLOUDINARY_API_KEY"),
		CloudinaryAPISecret: os.Getenv("CLOUDINARY_API_SECRET"),

		EmailExpires: os.Getenv("EMAIL_EXPIRES"),
		EmailHost:    os.Getenv("EMAIL_HOST"),
		EmailPort:    getEnv("EMAIL_PORT", "587"),
		EmailAddress: os.Getenv("EMAIL_ADDRESS"),
		EmailPass:    os.Getenv("EMAIL_PASS"),
		EmailFrom:    os.Getenv("EMAIL_FROM"),
		AdminEmail:   os.Getenv("ADMIN_EMAIL"),
	}

	if config.DSN == "" {
		return nil, fmt.Errorf("DSN is required")
	}

	if config.CloudinaryCloudName == "" {
		return nil, fmt.Errorf("CLOUDINARY_CLOUD_NAME is required")
	}

	if config.CloudinaryAPIKey == "" {
		return nil, fmt.Errorf("CLOUDINARY_API_KEY is required")
	}

	if config.CloudinaryAPISecret == "" {
		return nil, fmt.Errorf("CLOUDINARY_API_SECRET is required")
	}

	if config.EmailHost == "" {
		return nil, fmt.Errorf("EMAIL_HOST is required")
	}

	if config.EmailAddress == "" {
		return nil, fmt.Errorf("EMAIL_ADDRESS is required")
	}

	if config.EmailPass == "" {
		return nil, fmt.Errorf("EMAIL_PASS is required")
	}

	emailPort, err := strconv.Atoi(config.EmailPort)
	if err != nil || emailPort < 1 || emailPort > 65535 {
		return nil, fmt.Errorf("EMAIL_PORT must be a number between 1 and 65535")
	}

	return config, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}
