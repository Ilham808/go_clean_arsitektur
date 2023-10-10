package user

import (
	"GoClean/domain/user"
	"GoClean/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	usecase       user.UserUsecase
	secret        string
	refreshSecret string
}

func NewUserHandler(usecase user.UserUsecase, secret, refreshSecret string) *UserHandler {
	return &UserHandler{
		usecase,
		secret,
		refreshSecret,
	}
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

func (h *UserHandler) Login(c echo.Context) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	user, err := h.usecase.Authenticate(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"message": "Invalid username or password",
		})
	}

	token, err := helper.GenerateToken(h.secret, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	refreshToken, err := helper.GenerateRefreshToken(h.refreshSecret, user.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Login Successful",
		"access_token":  token,
		"refresh_token": refreshToken,
	})
}
