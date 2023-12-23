package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oliveirabalsa/balsacar-be/internal/config"
	"github.com/oliveirabalsa/balsacar-be/internal/middleware"
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
	secretKey := os.Getenv("SECRET_KEY")

	if port == "" {
		port = "8080"
	}

	advertisementHandler := config.AdvertisementHandlerFactory(db)
	authHandler := config.AuthHandlerFactory(db, []byte(secretKey))
	authMiddleware := middleware.AuthMiddleware([]byte(secretKey))

	router.InitRouter(server, advertisementHandler, authHandler, authMiddleware)

	server.Run(fmt.Sprintf(":%s", port))
}
