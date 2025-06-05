package main

import (
	"backend-go-demo/internal/db"
	"backend-go-demo/internal/handler"
	"backend-go-demo/internal/middleware"
	"backend-go-demo/internal/repository"
	"backend-go-demo/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	//database, _ := db.NewSQLiteDB("users.db")
	database, _ := db.NewGormDB("demo.db")
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepo)
	userHandler := handler.NewUserHandler(authService)

	purchaseRepo := repository.NewPurchaseRepository(database)
	purchaseService := service.NewPurchaseService(purchaseRepo)
	purchaseHandler := handler.NewPurchaseHandler(purchaseService)

	r := gin.Default()
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)

	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/profile", userHandler.Profile)
	protected.POST("/purchase", purchaseHandler.Buy)
	protected.GET("/purchases", purchaseHandler.History)

	_ = r.Run(":8080")
}
