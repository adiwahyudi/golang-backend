package routes

import (
	"yukevent/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// Middleware
	e.GET("/user", controllers.GetUserController)
	e.GET("/user/:id", controllers.GetSpecificUserController)
	e.POST("/user", controllers.RegisterUserController)
	e.DELETE("/user/:id", controllers.DeleteSpecificUserByID)

	return e
}
