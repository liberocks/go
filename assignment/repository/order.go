package repository

var GET_ORDER_STATEMENT = `SELECT id, customer_name, ordered_at, created_at, updated_at FROM orders WHERE id = $1`
var GET_ORDERS_STATEMENT = `SELECT id, customer_name, ordered_at, created_at, updated_at FROM orders ORDER BY created_at DESC LIMIT $1 OFFSET $2`
var COUNT_ORDERS_STATEMENT = `SELECT COUNT(*) FROM orders`
var CREATE_ORDER_STATEMENT = `INSERT INTO orders (id, customer_name, ordered_at) VALUES ($1, $2, $3)`
var UPDATE_ORDER_STATEMENT = `UPDATE orders SET customer_name = $2, ordered_at = $3, updated_at = $4 WHERE id = $1`
var DELETE_ORDER_STATEMENT = `DELETE FROM orders WHERE id = $1`





