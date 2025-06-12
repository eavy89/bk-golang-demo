package main

import (
	"backend-go-demo/internal/config"
	"backend-go-demo/internal/db"
	"backend-go-demo/internal/handler"
	"backend-go-demo/internal/logger"
	"backend-go-demo/internal/middleware"
	"backend-go-demo/internal/repository"
	"backend-go-demo/internal/service"
	"context"
	"github.com/gin-gonic/gin"
)

func main() {
	log := logger.Get()

	database, err := db.NewGormDB(config.GetDBPath())
	if err != nil {
		//log.Fatalf("failed to connect database: %v", err)
		log.Fatal(context.Background(), "Failed to connect database",
			logger.Field{Key: "error", Value: err},
			logger.Field{Key: "component", Value: "database"},
		)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	authService := service.NewAuthService(userRepo)
	userHandler := handler.NewUserHandler(authService)
	log.Info(context.Background(), "User handler created")

	purchaseRepo := repository.NewPurchaseRepository(database)
	purchaseService := service.NewPurchaseService(purchaseRepo)
	purchaseHandler := handler.NewPurchaseHandler(purchaseService)
	log.Info(context.Background(), "Purchases handler created")

	r := gin.Default()
	r.POST("/login", userHandler.Login)
	r.POST("/register", userHandler.Register)
	log.Info(context.Background(), "Public routes created")

	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware())
	protected.GET("/profile", userHandler.Profile)
	protected.POST("/purchase", purchaseHandler.Buy)
	protected.GET("/purchases", purchaseHandler.History)
	log.Info(context.Background(), "Protected routes created")

	log.Info(context.Background(), "Starting server at 0.0.0.0:8080 address")
	_ = r.Run(config.GetServerPort())

}
