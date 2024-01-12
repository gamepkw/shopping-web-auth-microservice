package service

import (
	"context"
	"time"

	middleware "github.com/gamepkw/shopping-web-auth-microservice/internal/middleware"
	model "github.com/gamepkw/shopping-web-auth-microservice/internal/models"
	authRepo "github.com/gamepkw/shopping-web-auth-microservice/internal/repositories"
	utils "github.com/gamepkw/shopping-web-auth-microservice/internal/utils"
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
	ctx := context.Background()

	request.HashedPassword = utils.EncodeBase64(request.HashedPassword)

	if err := a.authRepo.Login(ctx, request); err != nil {
		return "", err
	}

	token, err := middleware.GenerateJWTToken(request.Username, 1*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}
