package service

import (
	"net/http"

	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func DeleteOrder(id string) (int, error) {
	db := helpers.GetDB()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/delete_order] Failed to start transaction: %v", err)
		return http.StatusInternalServerError, err
	}

	// Defer a rollback in case anything fails
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error().Err(err).Msgf("[repository/delete_order] Transaction rolled back due to panic: %v", r)
		}
	}()

	// Delete all items from the items table
	_, err = tx.Exec(repository.DELETE_ITEMS_STATEMENT, id)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[repository/delete_order] Failed to delete items: %v", err)

		return http.StatusNotFound, err
	}

	// Delete order from the orders table
	_, err = tx.Exec(repository.DELETE_ORDER_STATEMENT, id)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[repository/delete_order] Failed to update order: %v", err)

		return http.StatusNotFound, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/delete_order] Failed to commit transaction: %v", err)
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}
