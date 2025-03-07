package auth

import "go/adv-demo/internal/user"

type AuthService struct {
	UserRepository *user.UserRepository
}

func NewAuthService(userRepository *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepository: userRepository,
	}
}

// func (service *AuthService) Register(email, name, pasword string) (string, error){
	
// }