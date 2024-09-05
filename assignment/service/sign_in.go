package service

import (
	"net/http"

	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/model"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func SignIn(email string, password string) (dto.SignInResponse, int, error) {
	db := helpers.GetDB()
	var credentials = dto.SignInResponse{}

	// Check if the email already exists
	var user = model.User{}
	err := db.QueryRow(repository.GET_USER_STATEMENT, email).Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msgf("[repository/sign_in] Failed to get user: %v", err)

		return credentials, http.StatusNotFound, err
	}

	// Compare the password
	if err := helpers.ValidatePassword(password, user.Password); err != nil {
		log.Error().Err(err).Msgf("[repository/sign_in] Failed to validate password: %v", err)

		return credentials, http.StatusUnauthorized, err
	}

	// Create a JWT token
	token, err := helpers.CreateAccessToken(user.Id)

	credentials.AccessToken = token

	return credentials, http.StatusOK, nil
}
