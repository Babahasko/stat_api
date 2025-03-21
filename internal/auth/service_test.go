package auth_test

import (
	"github.com/Babahasko/stat_api/internal/auth"
	"github.com/Babahasko/stat_api/internal/user"
	"testing"
)

type MockUserRepository struct{}

func (repo *MockUserRepository) Create(u *user.User) (*user.User, error){
	return &user.User{
		Email: "a@a.ru",
	}, nil
}

func (repo *MockUserRepository) GetByEmail(email string) (*user.User, error){
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	const initialEmail = "a@a.ru"
	authService := auth.NewAuthService(&MockUserRepository{})
	email, err := authService.Register(initialEmail, "1", "Вася")
	if err != nil {
		t.Fatal(err)
	}
	if email != initialEmail {
		t.Fatalf("Expected %s not match with %s", initialEmail, email)
	}
}