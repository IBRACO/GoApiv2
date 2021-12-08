package controllers

import (
	// "fmt"
	"net/http"
	entity"example/fgp/entity"
	utils"example/fgp/utils"
	// controllers"example/fgp/controllers"
	// "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	database"example/fgp/database"
)

type data struct {
	Userdb  *entity.User
	Msg    string
}
type UserChangePassword struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	NewPassword string `json:"new_password"`
}

func ChangePassword(c echo.Context) error {
	db := database.Connexion()
	user := new(entity.User)
	userFront := new(UserChangePassword)
	if err := c.Bind(userFront); err != nil {
		return err
	}
	result := db.Where("email = ?", userFront.Email).First(&user)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "bad email adresse")
	}
	match, err := utils.ComparePassword(userFront.Password, user.Password)
	if !match || err != nil {
		return c.String(http.StatusBadRequest, "bad password")
	}

	hash, err := utils.GeneratePassword(userFront.NewPassword)
	if err != nil {
		panic(err)
	}

	user.Password = hash
	db.Save(&user)

	
	return c.JSON(http.StatusOK,data{user,"Vous avez bien changer votre mot de pass"+user.First_name})
}
