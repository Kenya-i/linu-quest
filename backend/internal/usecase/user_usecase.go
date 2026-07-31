package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Kenya-i/linu-quest/internal/domain"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo  domain.UserRepository
	jwtSecret []byte
}

func NewUserUsecase(userRepo domain.UserRepository, jwtSecret []byte) domain.UserUsecase {
	return &userUsecase{userRepo: userRepo, jwtSecret: jwtSecret}
}

func (u *userUsecase) Register(ctx context.Context, username string, email string, password string) (*domain.User, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username:       username,
		Email:          email,
		HashedPassword: string(hashed),
	}

	if err := u.userRepo.SaveUser(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

var ErrInvalidCredentials = errors.New("invalid username or password")

func (u *userUsecase) Login(ctx context.Context, username string, password string) (string, error) {
	user, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return "", ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password)); err != nil {
		return "", ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(u.jwtSecret)

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
