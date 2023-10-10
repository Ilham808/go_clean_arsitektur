package mocks

import (
	"GoClean/domain/user"

	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Create(user *user.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserRepositoryMock) FindAll() ([]user.User, error) {
	args := m.Called()
	return args.Get(0).([]user.User), args.Error(1)
}

func (m *UserRepositoryMock) Authenticate(email, password string) (*user.User, error) {
	args := m.Called(email, password)
	return args.Get(0).(*user.User), args.Error(1)
}
