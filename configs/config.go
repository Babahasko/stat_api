package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB DBConfig
}

type DBConfig struct{
	Dsn string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Erro loading .env file, using default config")
	}
	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
	}
}