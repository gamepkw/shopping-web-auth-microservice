package service

import (
	"context"

	model "github.com/gamepkw/shopping-web-auth-microservice/internal/models"
	authRepo "github.com/gamepkw/shopping-web-auth-microservice/internal/repositories"
)

type authService struct {
	authRepo authRepo.AuthRepository
}

func NewAuthService(ur authRepo.AuthRepository) *authService {
	return &authService{
		authRepo: ur,
	}
}

type AuthService interface {
	Register(c context.Context, request model.RegisterRequest) error
	Login(c context.Context, request model.LoginRequest) (string, error)
}

func (a *authService) Register(c context.Context, request model.RegisterRequest) error {
	// ctx := context.Background()

	return nil
}

func (a *authService) Login(c context.Context, request model.LoginRequest) (string, error) {
	// ctx := context.Background()

	return "", nil
}
