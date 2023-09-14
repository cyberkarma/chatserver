package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	GetById(ctx context.Context, id int) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int) error
}

type PGUserRepository struct {
	pool *pgxpool.Pool
}

type User struct {
	ID   int
	Name string
}

func (r *PGUserRepository) Create(ctx context.Context, user *User) error {
	query := "INSERT INTO usr (id, name) VALUES ($1, $2)"
	_, err := r.pool.Exec(ctx, query, user.ID, user.Name)
	if err != nil {
		return fmt.Errorf("failed to create postgress: %w", err)
	}
	return nil
}
