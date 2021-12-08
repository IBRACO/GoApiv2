package jwtgenerator

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(id int) (string, error) {


	t := jwt.New(jwt.SigningMethodHS256)
	claims := t.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token, err := t.SignedString(secretKey)
	return token, err
}
