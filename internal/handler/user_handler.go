package handler

import (
	"backend-go-demo/internal/config"
	"net/http"

	"backend-go-demo/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.AuthService
}

func (uh *UserHandler) Register(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := uh.service.Register(req.Username, req.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "registered"})
}

func (uh *UserHandler) Login(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	token, err := uh.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (uh *UserHandler) Profile(c *gin.Context) {
	userID := c.GetInt(config.ClaimUserID)
	username := c.GetString(config.ClaimUsername)
	c.JSON(http.StatusOK, gin.H{"user": username, "id": userID})
}

func NewUserHandler(service *service.AuthService) *UserHandler {
	return &UserHandler{service: service}
}
