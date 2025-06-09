package db

import (
	"backend-go-demo/internal/model"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type GormDB struct {
	conn *gorm.DB
}

func (g *GormDB) CreateUser(user *model.User) error {
	return g.conn.Create(user).Error
}

func (g *GormDB) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := g.conn.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (g *GormDB) CreatePurchase(purchase *model.Purchase) error {
	purchase.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	return g.conn.Create(purchase).Error
}

func (g *GormDB) GetPurchasesByUserID(userID int) ([]model.Purchase, error) {
	var purchases []model.Purchase
	err := g.conn.Where("user_id = ?", userID).Find(&purchases).Error
	return purchases, err
}

func (g *GormDB) Close() error {
	dbSQL, err := g.conn.DB()
	if err != nil {
		return err
	}
	return dbSQL.Close()
}

func NewGormDB(path string) (*GormDB, error) {
	// Ensure the folder exists
	if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create db directory: %w", err)
	}

	conn, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open db: %w", err)
	}

	if err := conn.AutoMigrate(&model.User{}, &model.Purchase{}); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	return &GormDB{conn: conn}, nil
}
