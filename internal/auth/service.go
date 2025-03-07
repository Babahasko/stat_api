package auth

import (
	"errors"
	"go/adv-demo/internal/user"
)

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

func (service *AuthService) Register(email, name, pasword string) (string, error) {
	existed_user, _ := service.UserRepository.GetByEmail(email)
	if existed_user != nil {
		return "", errors.New(ErrorUserExists)
	}
	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	_, err := service.UserRepository.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
