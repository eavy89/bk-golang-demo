package repository

import (
	"backend-go-demo/internal/db"
	"backend-go-demo/internal/model"

	_ "github.com/mattn/go-sqlite3"
)

type UserRepository struct {
	db db.DB
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.db.CreateUser(user)
}

func (ur *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	return ur.db.GetUserByUsername(username)
}

func NewUserRepository(db db.DB) *UserRepository {
	return &UserRepository{db: db}
}
