package connection

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	DB *sqlx.DB
}

// NewDatabase is
func NewDatabase() (*Database, error) {
	db, err := sqlx.Open("mysql", "root:NewPassword!123@tcp(localhost:3306)/ecomm_mysql?parseTime=true")
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return &Database{DB: db}, nil
}

func (d *Database) Close() error {
	return d.DB.Close()
}

func (d *Database) GetDB() *sqlx.DB {
	return d.DB
}
