package jwtgenerator

import (
	// jwtgenerator"example/fgp/jwtgenerator"
	"fmt"
	"net/http"

	// jwt"example/fgp/jwt"
	"github.com/dgrijalva/jwt-go"

	"github.com/labstack/echo"
	utils "example/fgp/utils"
)


var secretKey  string
func Token(c echo.Context) error {

	secretKey,_ = utils.GenerateRsa()
	type tokenReqBody struct {
		RefreshToken string `json:"refresh_token"`
	}
	tokenReq := tokenReqBody{}
	// c.Bind(&tokenReq)
	if err := c.Bind(tokenReq); err != nil {
		return err
	}

	token, err := jwt.Parse(tokenReq.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if int(claims["iss"].(float64)) == 1 {

			newTokenPair, err := GenerateTokenPair()
			if err != nil {
				return err
			}

			return c.JSON(http.StatusOK, newTokenPair)
		}

		return echo.ErrUnauthorized
	}

	return err
}