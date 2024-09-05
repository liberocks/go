package dto

type GetOrderDetailItemResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type GetOrderDetailResponse struct {
	Id           string                       `json:"id"`
	OrderedAt    string                       `json:"orderedAt"`
	CreatedAt    string                       `json:"createdAt"`
	UpdatedAt    string                       `json:"updatedAt"`
	CustomerName string                       `json:"customerName"`
	Items        []GetOrderDetailItemResponse `json:"items"`
}
