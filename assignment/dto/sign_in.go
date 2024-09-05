package dto

import (
	"github.com/liberocks/go/assignment/helpers"
)

type SignInPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type SignInResponse struct {
	AccessToken  string `json:"accessToken"`
}

func (d *SignInPayload) Validate() error {
	validator := helpers.GetValidator()
	return validator.Struct(d)
}
