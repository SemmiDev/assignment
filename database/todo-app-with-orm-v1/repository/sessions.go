package repository

import (
	"a21hc3NpZ25tZW50/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SessionsRepository struct {
	db *gorm.DB
}

func NewSessionsRepository(db *gorm.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	return u.db.Create(&session).Error
}

func (u *SessionsRepository) DeleteSession(token string) error {
	return u.db.Where("token = ?", token).Delete(&model.Session{}).Error
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	// update where username
	return u.db.Model(&model.Session{}).Where("username = ?", session.Username).Update("token", session.Token).Error
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	var session model.Session
	err := u.db.Where("username = ?", name).First(&session).Error
	return session, err
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	var session model.Session
	err := u.db.Where("token = ?", token).First(&session).Error
	return session, err
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}

	return session, nil
}

func (u *SessionsRepository) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}
