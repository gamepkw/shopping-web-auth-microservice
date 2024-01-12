package repository

import (
	"context"
	"fmt"

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
	Register(ctx context.Context, u model.RegisterMysql) error
	Login(ctx context.Context, u model.LoginRequest) error
}

type userCredential struct {
	Username       string `gorm:"username"`
	HashedPassword string `gorm:"hashed_password"`
}

func (m *authRepository) Register(ctx context.Context, request model.RegisterMysql) error {

	var userCredential userCredential

	if err := m.db.Table("user_credentials").Where("username = ?", request.Username).First(&userCredential).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	if len(userCredential.Username) != 0 {
		return fmt.Errorf("username exists")
	}

	if err := m.db.Table("user_credentials").Create(&request).Error; err != nil {
		return fmt.Errorf("an error occurred while creating user: %w", err)
	}

	return nil
}

func (m *authRepository) Login(ctx context.Context, request model.LoginRequest) error {

	hashPassword, err := m.getHashedPassword(ctx, request.Username)
	if err != nil {
		return err
	}

	if request.HashedPassword != hashPassword {
		return fmt.Errorf("wrong password")
	}

	return nil
}

func (m *authRepository) getHashedPassword(ctx context.Context, username string) (string, error) {

	var userCredential userCredential
	if err := m.db.Table("user_credentials").Where("username = ?", username).First(&userCredential).Error; err != nil {
		return "", err
	}

	return userCredential.HashedPassword, nil
}
