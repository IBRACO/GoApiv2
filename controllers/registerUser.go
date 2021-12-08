package controllers

import (

	"net/http"
	entity"example/fgp/entity"
	utils"example/fgp/utils"
	database"example/fgp/database"

	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	db := database.Connexion()
	// db := connexion()
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	// password hash before saving

	password := user.Password
	emailSended := user.Email
	// db.First(&user, emailSended)
	// user2 := new(User)
	result1 := db.Where("email = ?",emailSended ).First(&user)

	
	if result1.Error != nil {

		 hash, _ := utils.GeneratePassword(password)
	     user.Password = hash
		result := db.Create(&user)
		if result.RowsAffected == 0 {
			panic(result.Error)

			// eror := Er{result.Error, "yu have to use a other email adresse"}

			// return c.JSON(http.StatusBadRequest, eror)

			// fmt.Sprintf("IMPOSIBLE DE CREER %v", result.Error)
			// return c.String(http.StatusBadRequest, fmt.Sprintf("IMPOSIBLE DE CREER\n  %v", result.Error))
		}

		return c.JSON(http.StatusOK,  user)

	}
	return c.String(http.StatusBadRequest,"Adresse  mail occupé par une autre personne, veuiller utiliser une autre adresse \n ou vous autentifier")
}
