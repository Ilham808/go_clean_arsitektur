package mocks

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type UserHandlerMock struct {
	mock.Mock
}

func (m *UserHandlerMock) Create(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *UserHandlerMock) FindAll(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *UserHandlerMock) Login(c echo.Context) error {
	args := m.Called(c)
	return args.Error(0)
}
