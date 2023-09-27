package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
	"github.com/oliveirabalsa/balsacar-be/internal/repository"
	"github.com/oliveirabalsa/balsacar-be/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	router := gin.Default()
	err := godotenv.Load()

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.Table("advertisements").AutoMigrate(&entity.Advertisement{})

	advertisementRepository := repository.NewAdvertisementRepository(db)
	advertisementService := service.NewAdvertisementService(advertisementRepository)
	advertisementHandler := handler.NewAdvertisementHandler(advertisementService)

	api := router.Group("/api")
	{
		api.POST("/advertisements", advertisementHandler.CreateAdvertisementHandler)
		api.GET("/advertisements/:id", advertisementHandler.GetAdvertisementByIDHandler)
		api.GET("/advertisements", advertisementHandler.GetAllAdvertisementsHandler)
		api.PUT("/advertisements/:id", advertisementHandler.UpdateAdvertisementHandler)
		api.DELETE("/advertisements/:id", advertisementHandler.DeleteAdvertisementHandler)
		api.POST("/advertisements/upload", advertisementHandler.UploadSheetAdvertisementHandler)
	}

	router.Run(":8080")
}
