package database

import (
	"GoClean/domain/user"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) Create(user *user.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepositoryImpl) FindAll() ([]user.User, error) {
	var users []user.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
