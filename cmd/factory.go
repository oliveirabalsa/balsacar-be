package main

import (
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
	"github.com/oliveirabalsa/balsacar-be/internal/repository"
	"github.com/oliveirabalsa/balsacar-be/internal/service"
	"gorm.io/gorm"
)

func AdvertisementHandlerFactory(db *gorm.DB) *handler.AdvertisementHandler {
	advertisementRepository := repository.NewAdvertisementRepository(db)
	advertisementService := service.NewAdvertisementService(advertisementRepository)
	return handler.NewAdvertisementHandler(advertisementService)
}

func AuthHandlerFactory(db *gorm.DB) *handler.AuthenticationHandler {
	authRepository := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepository, []byte("12345678"))
	return handler.NewAuthHandler(authService)
}
