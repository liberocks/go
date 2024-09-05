package service

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/model"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func SignUp(email string, password string) (dto.SignUpResponse, int, error) {
	db := helpers.GetDB()

	var response = dto.SignUpResponse{}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/sign_up] Failed to start transaction: %v", err)
		return response, http.StatusInternalServerError, err
	}

	// Defer a rollback in case anything fails
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error().Err(err).Msgf("[repository/sign_up] Transaction rolled back due to panic: %v", r)
		}
	}()

	// Check if the email already exists
	var user = model.User{}
	tx.QueryRow(repository.GET_USER_STATEMENT, email).Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if user.Email != "" {
		tx.Rollback()
		log.Error().Msgf("[service/sign_up] Email already exists: %v", email)

		return response, http.StatusConflict, nil
	}

	// Hash the password
	hashedPassword, err := helpers.HashPassword(password)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[service/sign_up] Failed to hash password: %v", err)

		return response, http.StatusInternalServerError, err
	}

	// Generate uuid
	id, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Msg("[repository/create_order] Failed to generate UUID")
		return response, http.StatusInternalServerError, err
	}

	// Insert user into the users table
	err = tx.QueryRow(repository.INSERT_USER_STATEMENT, id, email, hashedPassword).Scan(&user.Id, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[service/sign_up] Failed to insert user: %v", err)

		return response, http.StatusInternalServerError, err
	}

	// Change user validation status into true for simplicity. In the real world, we should send an email to the user to verify their email address.
	_, err = tx.Exec(repository.UPDATE_USER_VALIDATION_STATEMENT, user.Id, true)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[service/sign_up] Failed to update user validation status: %v", err)

		return response, http.StatusInternalServerError, err
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/sign_up] Failed to commit transaction: %v", err)
		return response, http.StatusInternalServerError, err
	}

	response.UserId = user.Id

	return response, http.StatusCreated, nil
}
