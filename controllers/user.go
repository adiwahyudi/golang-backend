package controllers

import (
	"net/http"
	"strconv"
	"yukevent/lib/database"
	"yukevent/model"

	"github.com/labstack/echo/v4"
)

func GetAllUserController(c echo.Context) error {

	users, err := database.GetAllUser()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponseArray(*users))
}

func GetSpecificUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	user, err := database.GetOneUserByID(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponse(*user))
}

func RegisterUserController(c echo.Context) error {

	var userReq RequestUser

	c.Bind(&userReq)

	result, err := database.CreateUser(userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponse(*result))
}

func UpdateUserController(c echo.Context) error {

	var updateUserReq RequestUser

	c.Bind(&updateUserReq)

	id, _ := strconv.Atoi(c.Param("id"))

	result, err := database.UpdateUser(id, updateUserReq.toModel())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponse(*result))

}

func DeleteSpecificUserByIDController(c echo.Context) error {

	var user model.User

	id, _ := strconv.Atoi(c.Param("id"))
	_, errFind := database.GetOneUserByID(id)
	_, err := database.DeleteUser(id, user)
	if errFind != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "not found",
		})
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "deleted",
	})
}

func DeleteSoftController(c echo.Context) error {

	users, err := database.DeleteSoft()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponseArray(*users))

}
