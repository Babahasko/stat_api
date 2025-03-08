package auth

import (
	"errors"
	"go/adv-demo/internal/user"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func(service *AuthService) Login(email, password string) (string, error) {
	existed_user, _ := service.UserRepository.GetByEmail(email)
	if existed_user == nil {
		return "", errors.New(ErrorWrongCredentials)
	}

	err := bcrypt.CompareHashAndPassword([]byte(existed_user.Password), []byte(password))
	if err != nil {
		return "", errors.New(ErrorWrongCredentials)
	}
	return existed_user.Email, nil
}

func (service *AuthService) Register(email, name, password string) (string, error) {
	existed_user, _ := service.UserRepository.GetByEmail(email)
	if existed_user != nil {
		return "", errors.New(ErrorUserExists)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	
	user := &user.User{
		Email:    email,
		Password: string(hashedPassword),
		Name:     name,
	}
	_, err = service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
