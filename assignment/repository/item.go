package repository

var GET_ITEMS_STATEMENT = `SELECT id, name, description, quantity, created_at, updated_at FROM items WHERE order_id = $1`
var DELETE_ITEMS_STATEMENT = `DELETE FROM items WHERE order_id = $1`
var CREATE_ITEM_STATEMENT = `INSERT INTO items (id, name, description, quantity, order_id) VALUES ($1, $2, $3, $4, $5) RETURNING id`
