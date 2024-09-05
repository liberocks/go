package repository

var GET_USER_STATEMENT = `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
var INSERT_USER_STATEMENT = `INSERT INTO users (id, email, password) VALUES ($1, $2, $3) RETURNING id, email, password, created_at, updated_at`
var UPDATE_USER_VALIDATION_STATEMENT = `UPDATE users SET is_validated = $2 WHERE id = $1`
