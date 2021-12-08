package controllers

import (
	"net/http"
	"strings"
	entity"example/fgp/entity"
	utils"example/fgp/utils"
	database"example/fgp/database"
	jwtgenerator "example/fgp/jwtgenerator"
	

	// "time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	// db := connexion()
	db := database.Connexion()
	user := new(entity.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	passwordSend := user.Password
	// email := user.Email
	// user2 := User{}
	result := db.Where("email = ?", user.Email).First(&user)
	if result.Error != nil {
		return c.String(http.StatusBadRequest, "bad email adresse")
	}
	match,err := utils.ComparePassword(passwordSend,user.Password)
	if !match || err != nil {
		return c.String(http.StatusBadRequest, "bad password")
		 } 
		 	
	
		token, err := jwtgenerator.GenerateJWT(user.ID)
		if err != nil {
			return c.String(http.StatusBadRequest, "could not log in")
		}
		user.Token = token

		uuidWithHyphen := uuid.New()
	    uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	    user.UUID = uuid
		db.Save(&user)
		return c.JSON(http.StatusOK, user)

	}
	// Jwt creation de token
	// token, err := GenerateJWT(user.ID)
	// if err != nil {
	// 	return c.String(http.StatusBadRequest, "could not log in")
	// }
	// user.Token = token
	// db.Save(&user)

	// return c.JSON(http.StatusCreated, " conected successfully")
	



// set cookie
// cookie := new(http.Cookie)
// cookie.Name = "jwt"
// cookie.Value = token
// cookie.Expires = time.Now().Add(10 * time.Minute)
// cookie.HttpOnly = true
// // cookie.Secure = true
// c.SetCookie(cookie)
