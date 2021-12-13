package routs

import(
	"github.com/labstack/echo/v4"
	controllers"example/fgp/controllers"
) 


func Roots() {
	// db = connexion()
	e := echo.New()
	e.GET("/user/:id", controllers.GetUserById)
	e.PUT("/user/update/:id", controllers.UpdateUserById)
	e.DELETE("/user/delete/:id/:role", controllers.DeleteUserById)
	e.POST("/fg", controllers.ForgetPw)
	e.POST("/change", controllers.ChangePassword)
	e.POST("/uppass", controllers.UpdatePassword)
	e.POST("/save", controllers.SaveUser)
	e.POST("/login", controllers.Login)
	e.Logger.Fatal(e.Start(":1323"))
}


