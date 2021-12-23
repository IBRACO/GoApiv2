package controllers

import (
	database "example/fgp/database"
	entity "example/fgp/entity"
	// jwtgenerator "example/fgp/jwtgenerator"
	utils "example/fgp/utils"
	"net/http"
	"strings"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db := database.Connexion()
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	passwordSend := user.Password
	result := db.Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "bad email adresse")
	}
	match, err := utils.ComparePassword(passwordSend, user.Password)
	if !match || err != nil {
		return c.String(http.StatusBadRequest, "bad password")
	}

	// token, err := jwtgenerator.GenerateJWT(user.ID)
	// if err != nil {
	// 	return err
	// }
	// user.Token = token

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	user.UUID = uuid
	db.Save(&user)
	return c.JSON(http.StatusOK, user)

}

