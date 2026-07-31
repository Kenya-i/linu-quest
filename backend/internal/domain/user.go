package domain

import (
	"context"
	"time"
)

type User struct {
	ID             string
	Username       string
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
	SaveUser(ctx context.Context, user *User) error
}

type UserUsecase interface {
	Register(ctx context.Context, username string, email string, password string) (*User, error)
	Login(ctx context.Context, username string, password string) (string, error)
}
