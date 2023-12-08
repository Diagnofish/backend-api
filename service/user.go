package service

import (
	"diagnofish/model"
	repo "diagnofish/repository"
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserService interface {
	Register(user *model.User) (model.User, error)
	Login(user *model.User) (token *string, err error)
}

type userService struct {
	userRepo    repo.UserRepository
	sessionRepo repo.SessionRepository
}

func NewUserService(userRepository repo.UserRepository, sessionsRepo repo.SessionRepository) UserService {
	return &userService{userRepository, sessionsRepo}
}

func (s *userService) Register(user *model.User) (model.User, error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return *user, err
	}

	if dbUser.Email != "" || dbUser.ID != "" {
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

func (s *userService) Login(user *model.User) (token *string, err error) {
	dbUser, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if dbUser.Email == "" || dbUser.ID == "" {
		return nil, errors.New("user not found")
	}

	if user.Password != dbUser.Password {
		return nil, errors.New("wrong email or password")
	}

	expirationTime := time.Now().Add(200 * time.Minute)
	claims := &model.Claims{
		UserId: dbUser.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		return nil, err
	}

	session := model.Session{
		Token:  tokenString,
		UserId: user.ID,
		Expiry: expirationTime,
	}

	_, err = s.sessionRepo.SessionAvailUserId(session.UserId)
	if err != nil {
		err = s.sessionRepo.AddSession(session)
	} else {
		err = s.sessionRepo.UpdateSession(session)
	}

	return &tokenString, nil
}
