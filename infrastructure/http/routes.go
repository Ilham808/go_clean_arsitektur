package http

import (
	"GoClean/infrastructure/http/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetRoutesUser(e *echo.Echo, userHandler *user.UserHandler, secret, refreshSecret string) {
	e.POST("/users", userHandler.Create, echojwt.JWT([]byte(secret)))
	e.GET("/users", userHandler.GetAll, echojwt.JWT([]byte(secret)))
	e.POST("/login", userHandler.Login)
}
