package service

import (
	"net/http"

	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func GetOrderDetail(id string) (dto.GetOrderDetailResponse, int, error) {
	db := helpers.GetDB()

	var order = dto.GetOrderDetailResponse{}

	// Get order from the orders table
	err := db.QueryRow(repository.GET_ORDER_STATEMENT, id).Scan(&order.Id, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to get order: %v", err)

		return order, http.StatusNotFound, err
	}

	// Get items from the order_items table
	rows, err := db.Query(repository.GET_ITEMS_STATEMENT, id)
	if err != nil {
		log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to get items: %v", err)
		return order, http.StatusNotFound, err
	}

	// Iteratively scan items
	var items = []dto.GetOrderDetailItemResponse{}
	for rows.Next() {
		var item = dto.GetOrderDetailItemResponse{}
		err = rows.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to scan items: %v", err)
			return order, http.StatusNotFound, err
		}
		items = append(items, item)
	}

	order.Items = items

	return order, http.StatusOK, nil
}
