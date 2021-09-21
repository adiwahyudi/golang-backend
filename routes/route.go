package routes

import (
	"yukevent/controllers"
	m "yukevent/middlewares"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	e.GET("/user", controllers.GetAllUserController)
	m.LogMiddleware(e)
	e.GET("/user/:id", controllers.GetSpecificUserController)
	e.POST("/user", controllers.RegisterUserController)
	e.PUT("/user/:id", controllers.UpdateUserController)
	e.DELETE("/user/:id", controllers.DeleteSpecificUserByIDController)
	e.GET("/userDeleted", controllers.DeleteSoftController)

	eAuth := e.Group("/auth")
	eAuth.Use(echoMid.BasicAuth(m.BasicAuthDB))
	eAuth.GET("/user", controllers.GetAllUserController)

	return e
}
