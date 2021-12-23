package controllers

import (
	// "fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	jwtgenerator"example/fgp/jwtgenerator"
	
)


func Tk(c echo.Context) error {
tk, err :=  jwtgenerator.GenerateTokenPair()
if err != nil {
	panic(err)
}
	
	return c.JSON(http.StatusOK,tk)
}
