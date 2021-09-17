package controllers

import (
	"net/http"
	"yukevent/config"
	"yukevent/model"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []model.User

	err := config.DB.Find(&users).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get user success!",
		"data":    users,
	})
}

func RegisterUserController(c echo.Context) error {

	user := model.User{}

	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "create user success!",
		"data":    user,
	})
}
