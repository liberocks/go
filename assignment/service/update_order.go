package service

import (
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/model"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func UpdateOrder(id string, customerName string, orderedAt string, items []dto.UpdateOrderItemPayload) (string, int, error) {
	db := helpers.GetDB()

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/update_order] Failed to start transaction: %v", err)
		return "", http.StatusInternalServerError, err
	}

	// Defer a rollback in case anything fails
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Error().Err(err).Msgf("[repository/update_order] Transaction rolled back due to panic: %v", r)
		}
	}()

	// Check if the order exists
	var order = model.Order{}
	err = tx.QueryRow(repository.GET_ORDER_STATEMENT, id).Scan(&order.Id, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
	if err != nil || order.Id == "" {
		tx.Rollback()
		log.Error().Err(err).Msgf("[repository/update_order] Failed to get order: %v", err)

		return "", http.StatusNotFound, err
	}

	_, err = tx.Exec(repository.UPDATE_ORDER_STATEMENT, id, customerName, orderedAt, time.Now().UTC())
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[repository/update_order] Failed to update order: %v", err)

		return "", http.StatusNotFound, err
	}

	// Delete all items from the items table
	_, err = tx.Exec(repository.DELETE_ITEMS_STATEMENT, id)
	if err != nil {
		tx.Rollback()
		log.Error().Err(err).Msgf("[repository/update_order] Failed to delete items: %v", err)

		return "", http.StatusNotFound, err
	}

	// Insert items into the order_items table
	for _, item := range items {
		// Generate uuid
		itemId, err := uuid.NewV7()
		if err != nil {
			log.Error().Err(err).Msg("[repository/update_order] Failed to generate UUID")
			return "", http.StatusInternalServerError, err
		}

		_, err = tx.Exec(repository.CREATE_ITEM_STATEMENT, itemId, item.Name, item.Description, item.Quantity, id)
		if err != nil {
			tx.Rollback()
			log.Error().Err(err).Msgf("[repository/update_order] Failed to insert item: %v", err)

			return "", http.StatusNotFound, err
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Error().Err(err).Msgf("[repository/update_order] Failed to commit transaction: %v", err)
		return "", http.StatusInternalServerError, err
	}

	return id, http.StatusOK, nil
}
