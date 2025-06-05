package service

import (
	"backend-go-demo/internal/model"
	"backend-go-demo/internal/repository"
)

type PurchaseService struct {
	Repo *repository.PurchaseRepository
}

func (s *PurchaseService) Buy(p model.Purchase) error {
	return s.Repo.Save(p)
}

func (s *PurchaseService) History(userID int) ([]model.Purchase, error) {
	return s.Repo.FindByUserID(userID)
}

func NewPurchaseService(r *repository.PurchaseRepository) *PurchaseService {
	return &PurchaseService{Repo: r}
}
