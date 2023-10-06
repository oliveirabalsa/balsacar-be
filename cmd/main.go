package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oliveirabalsa/balsacar-be/internal/config"
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
	"github.com/oliveirabalsa/balsacar-be/internal/repository"
	"github.com/oliveirabalsa/balsacar-be/internal/service"
	"github.com/oliveirabalsa/balsacar-be/router"
)

func main() {
	server := gin.Default()
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	db := config.InitDB()
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	advertisementRepository := repository.NewAdvertisementRepository(db)
	advertisementService := service.NewAdvertisementService(advertisementRepository)
	advertisementHandler := handler.NewAdvertisementHandler(advertisementService)

	router.InitRouter(server, advertisementHandler)

	server.Run(fmt.Sprintf(":%s", port))
}
