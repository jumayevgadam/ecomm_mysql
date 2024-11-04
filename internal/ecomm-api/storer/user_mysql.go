package storer

import (
	"context"
	"fmt"

	"github.com/jumayevgadam/ecomm_mysql/internal/models"
)

// CreateUser method is
func (ms *MySQLStorer) CreateUser(ctx context.Context, u *models.User) (*models.User, error) {
	res, err := ms.DB.NamedExecContext(
		ctx,
		"INSERT INTO users (name, email, password, is_admin) VALUES (:name, :email, :password, :is_admin)",
		u,
	)
	if err != nil {
		return nil, fmt.Errorf("error inserting user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}
	u.ID = id

	return u, nil
}

// GetUser method is
func (ms *MySQLStorer) GetUser(ctx context.Context, email string) (*models.User, error) {
	var u models.User
	if err := ms.DB.GetContext(
		ctx,
		&u,
		"SELECT * FROM users WHERE email=?",
		email,
	); err != nil {
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &u, nil
}

// ListUsers method is
func (ms *MySQLStorer) ListUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	if err := ms.DB.SelectContext(
		ctx,
		&users,
		"SELECT * FROM users",
	); err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}

	return users, nil
}

// UpdateUser method is
func (ms *MySQLStorer) UpdateUser(ctx context.Context, u *models.User) (*models.User, error) {
	_, err := ms.DB.NamedExecContext(
		ctx,
		"UPDATE users SET name=:name, email:=email, password:=password, is_admin=:is_admin, updated_at=:updated_at WHERE id=:id",
		u,
	)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}

	return u, nil
}

// DeleteUser method is
func (ms *MySQLStorer) DeleteUser(ctx context.Context, id int64) error {
	_, err := ms.DB.ExecContext(
		ctx,
		"DELETE FROM users WHERE id=?",
		id,
	)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}

	return nil
}
