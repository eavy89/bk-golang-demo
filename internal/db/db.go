package db

import (
	"backend-go-demo/internal/model"
)

type DB interface {
	// User operations
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)

	// Purchase operations
	CreatePurchase(purchase *model.Purchase) error
	GetPurchasesByUserID(userID int) ([]model.Purchase, error)

	Close() error
}
