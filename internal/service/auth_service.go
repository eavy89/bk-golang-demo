package service

import (
	"backend-go-demo/internal/config"
	"backend-go-demo/internal/logger"
	jwtModel "backend-go-demo/internal/middleware"
	"context"
	"errors"
	"time"

	"backend-go-demo/internal/model"
	"backend-go-demo/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repository.UserRepository
}

func (as *AuthService) Register(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return as.repo.CreateUser(&model.User{Username: username, Password: string(hash)})
}

func (as *AuthService) Login(username, password string) (string, error) {
	log := logger.Get()

	user, err := as.repo.GetUserByUsername(username)
	if err != nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid password")
	}

	claims := jwtModel.JWTData{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret, err := config.GetJWTKey()
	if err != nil {
		log.Fatal(context.Background(), "Failed to get JWT key",
			logger.Field{Key: "error", Value: err},
			logger.Field{Key: "component", Value: "jwt"},
		)
		return "", err
	}
	return token.SignedString(secret)
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{repo: repo}
}
