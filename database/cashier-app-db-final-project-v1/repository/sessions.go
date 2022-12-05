package repository

import (
	"a21hc3NpZ25tZW50/model"
	"time"
	"fmt"

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

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	return u.db.Where("token = ?", tokenTarget).Delete(&model.Session{}).Error
}

func (u *SessionsRepository) UpdateSessions(session model.Session) error {
	return u.db.Model(&model.Session{}).Where("username = ?", session.Username).Updates(session).Error
}

func (u *SessionsRepository) TokenValidity(token string) (model.Session, error) {
	ses, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(ses) {
		return model.Session{}, gorm.ErrRecordNotFound
	}

	return ses, nil
}

func (u *SessionsRepository) SessionAvailName(name string) (model.Session, error) {
	ses := model.Session{}

	err := u.db.Where("username = ?", name).First(&ses).Error
	if err != nil {
		return model.Session{}, err
	}

	return ses, nil
}

func (u *SessionsRepository) SessionAvailToken(token string) (model.Session, error) {
	ses := model.Session{}

	err := u.db.Where("token = ?", token).First(&ses).Error
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(ses) {
		return model.Session{}, gorm.ErrRecordNotFound
	}

	return ses, nil
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	fmt.Println(s.Expiry)
	return s.Expiry.Before(time.Now())
}
