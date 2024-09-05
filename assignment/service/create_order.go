package service

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func CreateOrder(customerName string, orderedAt string, items []dto.CreateOrderItemPayload) (string, int, error) {
	db := helpers.GetDB()

	// generate uuid
	id, err := uuid.NewV7()
	if err != nil {
		log.Error().Err(err).Msg("[repository/create_order] Failed to generate UUID")
		return "", http.StatusInternalServerError, err
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/create_order] Failed to start transaction: %v", err)
		return "", http.StatusInternalServerError, err
	}

	// Defer a rollback in case anything fails
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error().Err(err).Msgf("[repository/create_order] Transaction rolled back due to panic: %v", r)
		}
	}()

	_, err = tx.Exec(repository.CREATE_ORDER_STATEMENT, id, customerName, orderedAt)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[repository/create_order] Failed to insert order: %v", err)

		return "", http.StatusInternalServerError, err
	}

	// Insert items into the order_items table
	for _, item := range items {
		_, err := tx.Exec(repository.CREATE_ITEM_STATEMENT, item.Name, item.Description, item.Quantity, id)
		if err != nil {
			tx.Rollback()
			log.Error().Err(err).Msgf("[repository/create_order] Failed to insert item: %v", err)

			return "", http.StatusInternalServerError, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/create_order] Failed to commit transaction: %v", err)
		return "", http.StatusInternalServerError, err
	}

	return id.String(), http.StatusCreated, nil
}
