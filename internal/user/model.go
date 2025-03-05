package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"index" json:"email"`
	Name string `json:"name"`
	Password string `json:"password"`
}

func NewUser(email string,name string, password string) *User{
	user := &User{
		Email: email,
		Name: name,
		Password: password,
	}
	return user
}