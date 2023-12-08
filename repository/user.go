package repository

import (
	"diagnofish/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	GetUserByEmail(email string) (model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	if result := r.db.Create(&user); result.Error != nil {
		return user, result.Error
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var user model.User

	if err := r.db.Model(&user).Select("*").Where("email = $1", email).Scan(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
