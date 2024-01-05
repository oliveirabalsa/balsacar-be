// @title           Swagger Example API
// @version         1.0
// @description     BalsaCar API
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.email  oliveirabalsa2@gmail.com

// @license.name  MTI
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8081
// @BasePath  /api

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
