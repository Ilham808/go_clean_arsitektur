package user

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUsecase interface {
	Create(user *User) error
	FindAll() ([]User, error)
	Authenticate(email, password string) (*User, error)
}

type userUsecase struct {
	userRepository UserRepository
}

func NewUserUsecase(userRepository UserRepository) UserUsecase {
	return &userUsecase{
		userRepository,
	}
}

func (u *userUsecase) Create(user *User) error {
	return u.userRepository.Create(user)
}

func (u *userUsecase) FindAll() ([]User, error) {
	return u.userRepository.FindAll()
}

func (u *userUsecase) Authenticate(email, password string) (*User, error) {
	user, err := u.userRepository.Authenticate(email, password)

	if err != nil {
		return nil, err
	}

	if user != nil && user.Password == password {
		return user, nil
	}

	return nil, errors.New("invalid email or password")
}
