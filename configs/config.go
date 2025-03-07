package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DB DBConfig
	Auth AuthConfig
}

type DBConfig struct{
	Dsn string
}

type AuthConfig struct{
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load(`D:\Programming\2_Learn\Go\adv-demo\.env`)
	if err != nil {
		log.Println("Erro loading .env file, using default config")
	}
	return &Config{
		DB: DBConfig{
			Dsn: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("TOKEN"),
		},
	}
}