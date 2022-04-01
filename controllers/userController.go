package controllers

import (
	"net/http"
	"transaksi/lib/database"
	responses "transaksi/lib/response"
	"transaksi/models"

	"github.com/labstack/echo/v4"
)

var user models.Users

//register user
func RegisterUsersController(c echo.Context) error {
	c.Bind(&user)
	duplicate, _ := database.GetUserByUsername(user.Username)
	if duplicate > 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Username was used, try another username",
		})
	}

	Password, _ := database.GeneratehashPassword(user.Password)
	user.Password = Password

	_, err := database.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}

	return c.JSON(http.StatusCreated, responses.StatusSuccessRegister(user.Username))

}
