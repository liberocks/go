package dto

import (
	"github.com/liberocks/go/assignment/helpers"
)

type GetOrdersQuery struct {
	Page  int `validate:"required,min=1"`
	Limit int `validate:"required,min=1,max=100"`
}

func (d *GetOrdersQuery) Validate() error {
	validator := helpers.GetValidator()
	return validator.Struct(d)
}

type GetOrdersDataItemResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type GetOrdersDataResponse struct {
	Id           string                       `json:"id"`
	OrderedAt    string                       `json:"orderedAt"`
	CreatedAt    string                       `json:"createdAt"`
	UpdatedAt    string                       `json:"updatedAt"`
	CustomerName string                       `json:"customerName"`
	Items        []GetOrdersDataItemResponse `json:"items"`
}

type GetOrdersResponse struct {
	Data      []GetOrdersDataResponse `json:"data"`
	Total     int                     `json:"total"`
	Page      int                     `json:"page"`
	Limit     int                     `json:"limit"`
	TotalPage int                     `json:"totalPage"`
}
