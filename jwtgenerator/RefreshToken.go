package jwtgenerator

import (
	// "strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
	utils "example/fgp/utils"
)

type MyCustomClaims struct {
	Name string `json:"name"`
	// Admin bool   `json:"admin"`
	jwt.StandardClaims
}

type MyCustomClaimsR struct {
	jwt.StandardClaims
}

func GenerateTokenPair() (map[string]string, error) {
	// Create token
	// id := strconv.Itoa(n)
	v, err := utils.GenerateRsa()
	if err != nil {
		panic(err)
	}
	// id := string(id)
	var SECRET_KEY = []byte(v)
	//Acces token claims
	claims := MyCustomClaims{
		"JWT",
		// true,

		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:   "1",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return nil, err
	}
	//Refresh token claims
	claimsR := MyCustomClaimsR{

		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:   "1",
		},
	}

	tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsR)
	tr, err := tokenR.SignedString(SECRET_KEY)
	if err != nil {
		return nil, err
	}
	tk := map[string]string{
		"token":        t,
		"refreshToken": tr,
	}

	return tk, nil
}

