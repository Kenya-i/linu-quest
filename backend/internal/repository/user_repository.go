package repository

import (
	"context"

	"github.com/Kenya-i/linu-quest/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := `SELECT ID, USERNAME, EMAIL, HASHED_PASSWORD, CREATED_AT, UPDATED_AT FROM USERS WHERE USERNAME = $1`
	var user domain.User
	err := r.db.QueryRow(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.HashedPassword,
		&user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
