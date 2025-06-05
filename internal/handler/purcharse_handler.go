package handler

import (
	"backend-go-demo/internal/config"
	"backend-go-demo/internal/model"
	"backend-go-demo/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PurchaseHandler struct {
	Service *service.PurchaseService
}

func (ph *PurchaseHandler) Buy(c *gin.Context) {
	userID := c.GetInt(config.ClaimUserID) // set by JWT middleware

	var input struct {
		Item     string  `json:"item"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	fmt.Println("User ID: ", userID)

	purchase := model.Purchase{
		UserID:   userID,
		Item:     input.Item,
		Quantity: input.Quantity,
		Price:    input.Price,
	}

	err := ph.Service.Buy(purchase)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save purchase"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ph *PurchaseHandler) History(c *gin.Context) {
	userID := c.GetInt(config.ClaimUserID)

	purchases, err := ph.Service.History(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch history"})
		return
	}

	c.JSON(http.StatusOK, purchases)
}

func NewPurchaseHandler(s *service.PurchaseService) *PurchaseHandler {
	return &PurchaseHandler{Service: s}
}
