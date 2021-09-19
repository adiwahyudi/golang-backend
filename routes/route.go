package routes

import (
	"yukevent/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.GET("/user", controllers.GetAllUserController)
	e.GET("/user/:id", controllers.GetSpecificUserController)
	e.POST("/user", controllers.RegisterUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)
	e.DELETE("/user/:id", controllers.DeleteSpecificUserByIDController)
	e.GET("/userDeleted", controllers.DeleteSoftController)

	return e
}
