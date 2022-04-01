package routers

import (
	"net/http"
	"transaksi/constants"
	"transaksi/controllers"

	"github.com/labstack/echo/v4"
	echoMid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	//CORS
	e.Use(echoMid.CORSWithConfig(echoMid.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet,
			http.MethodPut,
			http.MethodPost,
			http.MethodDelete},
	}))

	// Midleware Auth
	r := e.Group("")
	r.Use(echoMid.JWT([]byte(constants.SECRET_JWT)))
	// ------------------------------------------------------------------
	// LOGIN & REGISTER USER
	// ------------------------------------------------------------------
	e.POST("/users", controllers.RegisterUsersController)

	// r.GET("/users/profile", controllers.GetUsersController)

	return e
}
