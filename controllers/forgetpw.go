package controllers

import (
	// "encoding/json"
	"fmt"
	// "log"
	"example/fgp/entity"
	"example/fgp/utils"
	"net/http"

	// "os/exec"
	"strings"
	// "strconv"
	database"example/fgp/database"
	"github.com/google/uuid"

	"github.com/labstack/echo/v4"
)


func ForgetPw(c echo.Context) error {
	db := database.Connexion()

	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	result1 := db.Where("email = ?", user.Email).First(&user)
	if result1.Error != nil {
		return c.String(http.StatusBadRequest, fmt.Sprintf("%v", result1.Error))
	}

	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	user.UUID = uuid
	// user.Token = token
	db.Save(&user)

	//send mail with token
	utils.Sendmail(user.Email, user.UUID)

	return c.JSON(http.StatusOK, user)
}

func UpdatePassword(c echo.Context) error {
	db := database.Connexion()

	user := new(entity.User)
	// user2 := new(User)
	if err := c.Bind(user); err != nil {
		return err
	}
	//  token := c.FormValue("token")
	passwordSended := user.Password

	result := db.Where("uuid = ?", user.UUID).First(&user)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("%v", result.Error))
	}

	hash, err := utils.GeneratePassword(passwordSended)
	if err != nil {
		panic(err)
	}

	user.Password = hash
	db.Save(&user)
	utils.Sendmail(user.Email, "vous avez changez votre mot de pass avec succee")

	return c.String(http.StatusOK, user.Password)
}
