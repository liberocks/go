package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/service"
)

func createOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	var payload dto.CreateOrderPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Validate request body
	if err := payload.Validate(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Creating orders
	orderId, status, err := service.CreateOrder(payload.CustomerName, payload.OrderedAt, payload.Items)
	if status != http.StatusCreated {
		http.Error(w, "", status)
		return
	} else if err != nil {

		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.CreateOrderResponse{OrderId: orderId})
}

func getOrdersHandler(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	rawQuery := r.URL.Query()

	query := dto.GetOrdersQuery{
		Page:  1,
		Limit: 10,
	}

	if rawQuery.Get("page") != "" {
		pageStr := rawQuery.Get("page")
		page, _ := strconv.Atoi(pageStr)
		query.Page = page
	}

	if rawQuery.Get("limit") != "" {
		limitStr := rawQuery.Get("limit")
		limit, _ := strconv.Atoi(limitStr)
		query.Limit = limit
	}

	// Validate query parameters
	if err := query.Validate(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	// Get orders using specified query parameters
	orders, status, err := service.GetOrders(query.Page, query.Limit)
	if status != http.StatusOK {
		http.Error(w, "", status)
		return
	} else if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orders)
}

func OrderHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getOrdersHandler(w, r)
		return
	} else if r.Method == "POST" {
		createOrderHandler(w, r)
		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
