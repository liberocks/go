package model

type Order struct {
	Id           string `db:"id"`
	CustomerName string `db:"customer_name"`
	OrderedAt    string `db:"ordered_at"`
	CreatedAt    string	`db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
	Items        []Item
}
