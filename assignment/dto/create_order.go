package dto

import (
	"github.com/liberocks/go/assignment/helpers"
)

type CreateOrderItemPayload struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=255"`
	Quantity    int    `json:"quantity" validate:"required,min=1,max=9999"`
}

type CreateOrderPayload struct {
	OrderedAt    string                   `json:"orderedAt" validate:"required"`
	CustomerName string                   `json:"customerName" validate:"required,max=255"`
	Items        []CreateOrderItemPayload `json:"items" validate:"required,dive,required"`
}

type CreateOrderResponse struct {
	OrderId string `json:"order_id"`
}

func (d *CreateOrderPayload) Validate() error {
	validator := helpers.GetValidator()
	return validator.Struct(d)
}
