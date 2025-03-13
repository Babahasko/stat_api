package main

import (
	"github.com/Babahasko/stat_api/internal/link"
	"github.com/Babahasko/stat_api/internal/stat"
	"github.com/Babahasko/stat_api/internal/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(`D:\Programming\2_Learn\Go\adv-demo\.env`)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")),  &gorm.Config{
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&link.Link{}, &user.User{}, &stat.Stat{})
	if err != nil {
        log.Printf("Ошибка при выполнении миграций: %v", err)
		return
    }
	log.Println("Миграции успешно выполнены")
}