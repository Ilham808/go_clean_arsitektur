package user

import (
	"errors"
	"testing"

	"GoClean/domain/user"
	mocks "GoClean/mocks/domain/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserUsecase_Create(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	userUsecase := user.NewUserUsecase(mockRepo)

	user := &user.User{
		Email:    "budiawanilham04@gmail.com",
		Password: "12345",
	}

	mockRepo.On("Create", user).Return(nil)

	err := userUsecase.Create(user)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Create", mock.AnythingOfType("*user.User"))
}

func TestUserUsecase_Create_Error(t *testing.T) {
	userRepoMock := new(mocks.UserRepositoryMock)
	userUsecase := user.NewUserUsecase(userRepoMock)

	user := &user.User{
		Email:    "budiawanilham04@gmail.com",
		Password: "12345",
	}

	expectedErr := errors.New("error creating user")
	userRepoMock.On("Create", user).Return(expectedErr)

	err := userUsecase.Create(user)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedErr.Error())
	userRepoMock.AssertExpectations(t)
}

func TestUserUsecase_FindAll(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	userUsecase := user.NewUserUsecase(mockRepo)

	mockUsers := []user.User{
		{
			Email:    "budiawanilham04@gmail.com",
			Password: "12345",
		},
		{
			Email:    "budiawanilham04@gmail.com",
			Password: "12345",
		},
	}
	mockRepo.On("FindAll").Return(mockUsers, nil)

	users, err := userUsecase.FindAll()

	assert.NoError(t, err)
	assert.Len(t, users, len(mockUsers))
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_Authenticate(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)

	userUsecase := user.NewUserUsecase(mockRepo)

	mockUser := &user.User{
		Email:    "budiawanilham04@gmail.com",
		Password: "12345",
	}
	mockRepo.On("Authenticate", "budiawanilham04@gmail.com", "12345").Return(mockUser, nil)

	authenticatedUser, err := userUsecase.Authenticate("budiawanilham04@gmail.com", "12345")

	assert.NoError(t, err)
	assert.NotNil(t, authenticatedUser)
	assert.Equal(t, mockUser, authenticatedUser)
	mockRepo.AssertExpectations(t)
}

func TestUserUsecase_AuthenticateInvalidCredentials(t *testing.T) {
	mockRepo := new(mocks.UserRepositoryMock)
	userUsecase := user.NewUserUsecase(mockRepo)

	mockUser := &user.User{
		Email:    "budiawanilham04@gmail.com",
		Password: "2323232",
	}

	mockRepo.On("Authenticate", mockUser.Email, mockUser.Password).Return(mockUser, errors.New("invalid credentials"))

	authenticatedUser, err := userUsecase.Authenticate(mockUser.Email, mockUser.Password)

	assert.Error(t, err)
	assert.Nil(t, authenticatedUser)
	mockRepo.AssertExpectations(t)
}
