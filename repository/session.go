package repository

import (
	"diagnofish/model"
	"time"

	"gorm.io/gorm"
)

type SessionRepository interface {
	AddSession(session model.Session) error
	DeleteSession(token string) error
	UpdateSession(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionRepo struct {
	db *gorm.DB
}

func NewSessionRepo(db *gorm.DB) *sessionRepo {
	return &sessionRepo{db}
}

func (s *sessionRepo) AddSession(session model.Session) error {
	if err := s.db.Create(&session).Error; err != nil {
		return err
	}

	return nil
}

func (s *sessionRepo) DeleteSession(token string) error {
	var session model.Session

	if err := s.db.Where("token = ?", token).Delete(&session).Error; err != nil {
		return err
	}

	return nil
}

func (s *sessionRepo) UpdateSession(session model.Session) error {
	email := session.Email

	if err := s.db.Model(&session).Where("email = ?", email).Updates(map[string]interface{}{
		"token":  session.Token,
		"email":  session.Email,
		"expiry": session.Expiry,
	}).Error; err != nil {
		return err
	}

	return nil
}

func (s *sessionRepo) SessionAvailEmail(email string) (model.Session, error) {
	var session model.Session

	if err := s.db.Where("email = ?", email).First(&session).Error; err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepo) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session

	if err := s.db.Where("token = ?", token).First(&session).Error; err != nil {
		return model.Session{}, err
	}

	return session, nil
}

func (s *sessionRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}

func (s *sessionRepo) TokenValidity(token string) (model.Session, error) {
	session, err := s.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if s.TokenExpired(session) {
		err := s.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}

		return model.Session{}, err
	}

	return session, nil
}
