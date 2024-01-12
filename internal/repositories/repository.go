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
	Register(ctx context.Context, u userCredential) error
	Login(ctx context.Context, u model.LoginRequest) error
}

type userCredential struct {
	Username       string `gorm:"username"`
	HashedPassword string `gorm:"hashed_password"`
}

func (m *authRepository) Register(ctx context.Context, request userCredential) error {

	var userCredential userCredential
	if err := m.db.Where("username = ?", request.Username).First(&userCredential).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return err
		}
	}

	if err := m.db.Create(&request).Error; err != nil {
		return fmt.Errorf("an error while creating user")
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

	var hashPassword string
	if err := m.db.Table("user_credentials").Where("username = ?", username).Select("hash_password").First(&hashPassword).Error; err != nil {
		return "", err
	}

	return hashPassword, nil
}
