package routs

import (
	controllers "example/fgp/controllers"

	"github.com/labstack/echo/v4"
)

func Roots() {
	// db = connexion()
	e := echo.New()
	e.GET("/user/:id", controllers.GetUserById)
	e.GET("/tk", controllers.Tk)
	e.PUT("/user/update/:id", controllers.UpdateUserById)
	e.DELETE("/user/delete/:id/:role", controllers.DeleteUserById)
	e.POST("/fg", controllers.ForgetPw)
	e.POST("/change", controllers.ChangePassword)
	e.POST("/uppass", controllers.UpdatePassword)
	e.POST("/save", controllers.SaveUser)
	e.POST("/login", controllers.Login)

	e.Logger.Fatal(e.Start(":1323"))
}
