package storer

import "github.com/jmoiron/sqlx"

// MySQLStorer is
type MySQLStorer struct {
	DB *sqlx.DB
}

// NewMySQLStorer is
func NewMySQLStorer(DB *sqlx.DB) *MySQLStorer {
	return &MySQLStorer{DB: DB}
}
