package http

import (
	"GoClean/infrastructure/http/user"

	"github.com/labstack/echo/v4"
)

func SetRoutesUser(e *echo.Echo, userHandler *user.UserHandler) {
	e.POST("/users", userHandler.Create)
	e.GET("/users", userHandler.GetAll)
}
