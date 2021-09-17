package controllers

import (
	"net/http"
	"strconv"
	"yukevent/config"
	"yukevent/model"

	"github.com/labstack/echo/v4"
)

// Find All User
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

// Find 1 Specific User
func GetSpecificUserController(c echo.Context) error {
	var user []model.User

	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.First(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get specific user success",
		"data":    user,
	})
}

// Create a User
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

// Delete Spesific User
func DeleteSpecificUserByID(c echo.Context) error {

	user := model.User{}

	id, _ := strconv.Atoi(c.Param("id"))

	err := config.DB.Delete(&user, id).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete succes!",
	})
}

// Update Spesific User

// func UpdateSpesificUser(c echo.Context) error {

// }
