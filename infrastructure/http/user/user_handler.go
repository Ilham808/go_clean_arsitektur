package user

import (
	"GoClean/domain/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase user.UserUsecase
}

func NewUserHandler(usecase user.UserUsecase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) Create(c echo.Context) error {
	u := &user.User{}
	if err := c.Bind(u); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid request",
			"error":   err.Error(),
		})
	}

	err := h.usecase.Create(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "error creating user",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new user",
		"data":    u,
	})
}

func (h *UserHandler) GetAll(c echo.Context) error {
	users, err := h.usecase.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed to get all users",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"data":    users,
	})
}
