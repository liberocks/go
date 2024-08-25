package repository

import (
	"errors"

	"github.com/liberocks/go/mini-challenge-5/model"
)

var session = []model.Session{}

type SessionRepository struct {}

func (SessionRepository) Create(d model.Session) {
	session=  append(session, d)
}

func (SessionRepository) GetBySession(s string) (model.Session, error) {
	var result *model.Session
	for _, d := range session {
		if d.Session == s {
			result = &d
			break
		}
		
	}

	if result == nil {
		return model.Session{}, errors.New("Session not found")
	}

	return *result, nil
}

func (SessionRepository) Delete(s string ) error {
	arrIndex := -1

	for i, d := range session {
		if d.Session == s {
			arrIndex = i
			break
		}
	}

	if arrIndex == -1 {
		return errors.New("Session not found")
	}

	session = append(session[:arrIndex], session[arrIndex+1:]...)

	return nil
}
