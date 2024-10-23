package storer

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (ms *MySQLStorer) ExecTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	// begin transaction
	tx, err := ms.DB.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error starting transaction: %w", err)
	}

	// callback function
	err = fn(tx)
	if err != nil {
		// rollback tx if error occur
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("error rolling back transaction: %w", err)
		}

		return fmt.Errorf("error in transaction: %w", err)
	}

	// commit transaction
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("error committing transaction: %w", err)
	}

	return nil
}
