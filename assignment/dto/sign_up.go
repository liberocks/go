package dto

import (
	"github.com/liberocks/go/assignment/helpers"
)

type SignUpPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (d *SignUpPayload) Validate() error {
	validator := helpers.GetValidator()
	return validator.Struct(d)
}

type SignUpResponse struct {
	UserId string `json:"id"`
}
