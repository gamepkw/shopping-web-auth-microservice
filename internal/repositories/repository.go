package repository

import (
	"context"

	model "github.com/gamepkw/shopping-web-auth-microservice/internal/models"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		db: db,
	}
}

type AuthRepository interface {
	Register(ctx context.Context, u model.LoginRequest) error
	Login(ctx context.Context, u model.LoginRequest) error
}

func (m *authRepository) Register(ctx context.Context, request model.LoginRequest) error {

	return nil
}

func (m *authRepository) Login(ctx context.Context, request model.LoginRequest) error {

	return nil
}
