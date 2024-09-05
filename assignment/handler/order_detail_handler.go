package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/liberocks/go/assignment/dto"
	"github.com/liberocks/go/assignment/service"
)

func updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var updateOrderPayload dto.UpdateOrderPayload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&updateOrderPayload)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := updateOrderPayload.Validate(); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	orderId, status, err := service.UpdateOrder(id, updateOrderPayload.CustomerName, updateOrderPayload.OrderedAt, updateOrderPayload.Items)
	if status != http.StatusOK {
		http.Error(w, "", status)
		return
	} else if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.UpdateOrderResponse{OrderId: orderId})
}

func deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	status, err := service.DeleteOrder(id)
	if status != http.StatusOK {
		http.Error(w, "", status)
		return
	} else if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dto.StatusResponse{Status: "ok"})
}

func getOrderDetailHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	order, status, err := service.GetOrderDetail(id)
	if status != http.StatusOK {
		http.Error(w, "", status)
		return
	}

	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)
}

func OrderDetailHandlers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getOrderDetailHandler(w, r)
		return
	} else if r.Method == "PUT" {
		updateOrderHandler(w, r)
		return
	} else if r.Method == "DELETE" {
		deleteOrderHandler(w, r)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
