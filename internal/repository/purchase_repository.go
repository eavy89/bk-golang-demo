package repository

import (
	"backend-go-demo/internal/db"
	"backend-go-demo/internal/model"
)

type PurchaseRepository struct {
	db db.DB
}

func (pr *PurchaseRepository) Save(p model.Purchase) error {
	return pr.db.CreatePurchase(&p)
}

func (pr *PurchaseRepository) FindByUserID(userID int) ([]model.Purchase, error) {
	return pr.db.GetPurchasesByUserID(userID)
}

func NewPurchaseRepository(db db.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}
