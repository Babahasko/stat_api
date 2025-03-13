package main

import (
	"bytes"
	"encoding/json"
	"go/adv-demo/internal/auth"
	"go/adv-demo/internal/user"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDb() *gorm.DB {
	err := godotenv.Load(`D:\Programming\2_Learn\Go\adv-demo\cmd\.env`)
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
func initData(db *gorm.DB) {
	db.Create(&user.User{
		Email:    "bob8@mail.ru",
		Password: "$2a$10$iY3dGkePOnVpYIHIo429lezOqW0hPNDjAaKnyblCbIO4EBQJq9D6q",
		Name:     "Alexey",
	})
}

func removeData(db *gorm.DB) {
	db.Unscoped().
	Where("email = ?", "bob8@mail.ru").
	Delete(&user.User{})
}

func TestSuccessLogin(t *testing.T) {
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "bob8@mail.ru",
		Password: "11114",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 200 {
		t.Fatalf("Expected %d got %d \n", 200, res.StatusCode)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}
	var resData auth.LoginResponse
	err = json.Unmarshal(body, &resData)
	if err != nil {
		t.Fatal(err)
	}
	if resData.Token == "" {
		t.Fatal("Token empty")
	}
	removeData(db)
}

func TestFailLogin(t *testing.T) {
	db := initDb()
	initData(db)
	ts := httptest.NewServer(App())
	defer ts.Close()

	data, _ := json.Marshal(&auth.LoginRequest{
		Email:    "bob7@mail.ru",
		Password: "11112",
	})
	res, err := http.Post(ts.URL+"/auth/login", "application/json", bytes.NewReader(data))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != 401 {
		t.Fatalf("Expected %d got %d \n", 401, res.StatusCode)
	}
	removeData(db)
}
