package repository

import (
	"errors"

	"github.com/liberocks/go/mini-challenge-5/model"
)

var user = []model.User{
	{Id: 1, Name: "John Doe", Email: "john.doe@example.com", Address: "123 Main St, Cityville, USA", Job: "Software Engineer", Reason: "I want to learn more about Go"},
}

type UserRepository struct{}

 

func (UserRepository) GetById(id int) (model.User, error) {
	var result *model.User
	for _, u := range user {
		if u.Id == id {
			result = &u
			break
		}
	}

	if result == nil {
		return model.User{}, errors.New("User not found")
	}

	return *result, nil
}

func (UserRepository) GetByEmail(email string) (model.User, error) {
	var result *model.User
	for _, u := range user {
		if u.Email == email {
			result = &u
			break
		}
	}

	if result == nil {
		return model.User{}, errors.New("User not found")
	}

	return *result, nil
}

 