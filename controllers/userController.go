package controllers

import "net/http"

var user models.Users

//register user
func RegisterUsersController(c echo.Context) error {
	c.Bind(&user)
	duplicate, _ := database.GetUserByEmail(user.Email)
	if duplicate > 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email was used, try another email",
		})
	}

	Password, _ := database.GeneratehashPassword(user.Password)
	user.Password = Password

	_, err := database.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
