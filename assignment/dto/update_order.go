package dto

import (
	"github.com/liberocks/go/assignment/helpers"
)

type UpdateOrderItemPayload struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"required,max=255"`
	Quantity    int    `json:"quantity" validate:"required,min=1,max=9999"`
}

type UpdateOrderPayload struct {
	OrderedAt    string                   `json:"orderedAt" validate:"required"`
	CustomerName string                   `json:"customerName" validate:"required,max=255"`
	Items        []UpdateOrderItemPayload `json:"items" validate:"required,dive,required"`
}

type UpdateOrderResponse struct {
	OrderId string `json:"id"`
}

func (d *UpdateOrderPayload) Validate() error {
	validator := helpers.GetValidator()
	return validator.Struct(d)
}
