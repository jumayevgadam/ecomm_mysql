package storer

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateSession method is
func (ms *MySQLStorer) CreateSession(ctx context.Context, s *models.Session) (*models.Session, error) {
	_, err := ms.DB.NamedExecContext(
		ctx,
		"INSERT INTO sessions (id, user_email, refresh_token, is_revoked, expires_at) VALUES (:id, :user_email, :refresh_token, :is_revoked, :expires_at)",
		s,
	)
	if err != nil {
		return nil, fmt.Errorf("error inserting session: %w", err)
	}

	return s, nil
}

// GetSession method is
func (ms *MySQLStorer) GetSession(ctx context.Context, id string) (*models.Session, error) {
	var s models.Session
	if err := ms.DB.GetContext(
		ctx,
		&s,
		"SELECT * FROM sessions WHERE id=?",
		id,
	); err != nil {
		return nil, fmt.Errorf("error getting session: %w", err)
	}

	return &s, nil
}

// RevokeSession method is
func (ms *MySQLStorer) RevokeSession(ctx context.Context, id string) error {
	_, err := ms.DB.NamedExecContext(
		ctx,
		"UPDATE sessions SET is_revoked=1 WHERE id=:id",
		map[string]interface{}{"id": id},
	)
	if err != nil {
		return fmt.Errorf("error revoking session: %w", err)
	}

	return nil
}

// DeleteSession method is
func (ms *MySQLStorer) DeleteSession(ctx context.Context, id string) error {
	_, err := ms.DB.ExecContext(
		ctx,
		"DELETE FROM sessions WHERE id=?",
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting session: %w", err)
	}

	return nil
}
