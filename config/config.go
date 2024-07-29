package config

import "os"

type Config struct{
	DBUser string
	DBPassword string
	DBHost string
	DBName string
	DBPort string
	APIKey string
}

func GetConfig() Config {
	return Config{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		APIKey:     os.Getenv("YOUTUBE_API_KEY"),
	}
}