package controllers
import (
	database "example/fgp/database"
	entity "example/fgp/entity"
	utils "example/fgp/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SaveUser(c echo.Context) error {
	db := database.Connexion()
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	password := user.Password
	emailSended := user.Email
	result1 := db.Where("email = ?", emailSended).First(&user)
	if result1.Error != nil {
		hash, _ := utils.GeneratePassword(password)
		user.Password = hash
		if user.Role == ""{
			user.Role = "client"
		}
		result := db.Create(&user)
		if result.RowsAffected == 0 {
			panic(result.Error)
		}
		return c.JSON(http.StatusOK, user)
	}
	return c.String(http.StatusBadRequest, "Adresse  mail occupé par une autre personne, veuiller utiliser une autre adresse \n ou vous autentifier")
}
