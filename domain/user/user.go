package user

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserUsecase interface {
	Create(user *User) error
	FindAll() ([]User, error)
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
