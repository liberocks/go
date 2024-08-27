package repository

import (
	"errors"

	"github.com/liberocks/go/mini-challenge-5/model"
)

var credential = []model.Credential{
	{UserId: 1, Password: "password"},
}

type CredentialRepository struct{}

 

func (CredentialRepository) GetById(id int) (model.Credential, error) {
	var result *model.Credential
	for _, d := range credential {
		if d.UserId == id {
			result = &d
			break
		}
	}

	if result == nil {
		return model.Credential{}, errors.New("Credential not found")
	}

	return *result, nil
}
 