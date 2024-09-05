package model

type Item struct {
	Id          string `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
	Quantity    int    `db:"quantity"`
	OrderId     string `db:"order_id"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
}
