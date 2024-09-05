package service

import (
	"net/http"

	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/helpers"
	"github.com/liberocks/go/assignment/repository"
	"github.com/rs/zerolog/log"
)

func GetOrders(page int, limit int) (dto.GetOrdersResponse, int, error) {
	db := helpers.GetDB()

	var orders = dto.GetOrdersResponse{}

	rows, err := db.Query(repository.GET_ORDERS_STATEMENT, limit, (page-1)*limit)
	if err != nil {
		log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to get order: %v", err)

		return orders, http.StatusInternalServerError, err
	}

	for rows.Next() {
		var order = dto.GetOrdersDataResponse{}
		err = rows.Scan(&order.Id, &order.CustomerName, &order.OrderedAt, &order.CreatedAt, &order.UpdatedAt)
		if err != nil {
			log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to scan order: %v", err)

			return orders, http.StatusInternalServerError, err
		}

		// Insert items into the order_items table
		itemRows, err := db.Query(repository.GET_ITEMS_STATEMENT, order.Id)
		if err != nil {
			log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to get items: %v", err)
			return orders, http.StatusNotFound, err
		}

		var items = []dto.GetOrdersDataItemResponse{}
		for itemRows.Next() {
			var item = dto.GetOrdersDataItemResponse{}
			err = itemRows.Scan(&item.Id, &item.Name, &item.Description, &item.Quantity, &item.CreatedAt, &item.UpdatedAt)
			if err != nil {
				log.Error().Err(err).Msgf("[repository/get_order_detail] Failed to scan items: %v", err)
				return orders, http.StatusNotFound, err
			}
			items = append(items, item)
		}
		order.Items = items

		orders.Data = append(orders.Data, order)
	}

	// count total orders
	var totalOrders int
	err = db.QueryRow(repository.COUNT_ORDERS_STATEMENT).Scan(&totalOrders)

	orders.Limit = limit
	orders.Page = page
	orders.Total = totalOrders
	orders.TotalPage = (totalOrders + limit - 1) / limit

	return orders, http.StatusOK, nil
}
