package repository

import (
	"context"
	"database/sql"

	"github.com/TheAditya-10/GO/phase5/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {

	query := `
		SELECT id, tenant_id, email, password_hash, role, created_at
		FROM users
		WHERE email = ?
	`

	row := r.db.QueryRowContext(ctx, query, email)

	var u models.User
	err := row.Scan(
		&u.ID,
		&u.TenantID,
		&u.Email,
		&u.Password,
		&u.Role,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
