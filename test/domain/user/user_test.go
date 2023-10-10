package user

import (
	"testing"

	"GoClean/domain/user"
	mocks "GoClean/mocks/domain/user"

	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_Create(t *testing.T) {
	userRepoMock := new(mocks.UserRepositoryMock)
	userUsecase := user.NewUserUsecase(userRepoMock)

	user := &user.User{
		Email:    "budiawanilham04@gmail.com",
		Password: "12345",
	}

	userRepoMock.On("Create", user).Return(nil)

	err := userUsecase.Create(user)

	assert.NoError(t, err)
	userRepoMock.AssertExpectations(t)
}
