package repository

import (
	"a21hc3NpZ25tZW50/db"
	"a21hc3NpZ25tZW50/model"
	"encoding/json"
	"fmt"
	"time"
)

type SessionsRepository struct {
	db db.DB
}

func NewSessionsRepository(db db.DB) SessionsRepository {
	return SessionsRepository{db}
}

func (u *SessionsRepository) ReadSessions() ([]model.Session, error) {
	records, err := u.db.Load("sessions")
	if err != nil {
		return nil, err
	}

	var listSessions []model.Session
	err = json.Unmarshal([]byte(records), &listSessions)
	if err != nil {
		return nil, err
	}

	return listSessions, nil
}

func (u *SessionsRepository) DeleteSessions(tokenTarget string) error {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	if len(listSessions) == 1 {
		listSessions = []model.Session{}
	} else {
		for i, session := range listSessions {
			if session.Token == tokenTarget {
				listSessions = append(listSessions[:i], listSessions[i+1:]...)
				break
			}
		}
	}

	sessionData, err := json.Marshal(listSessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", sessionData)
	return err
}

func (u *SessionsRepository) AddSessions(session model.Session) error {
	sessions, err := u.ReadSessions()
	if err != nil {
		return err
	}

	sessions = append(sessions, session)
	sessionsData, err := json.Marshal(sessions)
	if err != nil {
		return err
	}

	err = u.db.Save("sessions", sessionsData)
	return err
}

func (u *SessionsRepository) CheckExpireToken(token string) (model.Session, error) {
	session, err := u.TokenExist(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		return model.Session{}, fmt.Errorf("Token is Expired!")
	}
	return session, nil
}

func (u *SessionsRepository) ResetSessions() error {
	err := u.db.Reset("sessions", []byte("[]"))
	if err != nil {
		return err
	}

	return nil
}

func (u *SessionsRepository) TokenExist(req string) (model.Session, error) {
	listSessions, err := u.ReadSessions()
	if err != nil {
		return model.Session{}, err
	}
	for _, element := range listSessions {
		if element.Token == req {
			return element, nil
		}
	}
	return model.Session{}, fmt.Errorf("Token Not Found!")
}

func (u *SessionsRepository) TokenExpired(s model.Session) bool {
	return s.Expiry.Before(time.Now())
}
