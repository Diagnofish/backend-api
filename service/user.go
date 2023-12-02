package service

import (
	"diagnofish/model"
	"diagnofish/repository"
	"errors"
	"strings"
	"time"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != 0 {
		return *user, errors.New("email has already exists")
	}

	username := strings.Split(user.Email, "@")

	user.Username = username[0]
	user.CreatedAt = time.Now()

	newUser, err := s.userRepo.CreateUser(*user)
	if err != nil {
		return *user, err
	}

	return newUser, nil
}
