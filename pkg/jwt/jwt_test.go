package jwt_test

import (
	"go/adv-demo/pkg/jwt"
	"testing"
)

func TestJWTCreate(t *testing.T) {
	const email = "mbator@yandex.ru"
	jwtService := jwt.NewJWT("xYKyuaQNvyQoJEhcWsB2sI5831lw2mTety3YiJO53k0=")
	jwtToken, err := jwtService.Create(&jwt.JWTData{
		Email: email,
	})
	if err != nil {
		t.Fatal(err)
	}
	isValid, data := jwtService.Parse(jwtToken)
	if !isValid {
		t.Fatal("token is Invalid")
	}
	if data.Email != email {
		t.Fatalf("Email %s not equal %s", data.Email, email)
	}


}
