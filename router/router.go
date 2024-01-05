package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/balsacar-be/docs"
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

func InitRouter(router *gin.Engine, advertisementHandler *handler.AdvertisementHandler, authenticationHandler *handler.AuthenticationHandler, authMiddleware gin.HandlerFunc) {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "balsacar.swagger.io"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		// Advertisement routes
		advertisements := api.Group("/advertisements", authMiddleware) // Change router to api here
		{
			advertisements.POST("/", advertisementHandler.CreateAdvertisementHandler)
			advertisements.GET("/:id", advertisementHandler.GetAdvertisementByIDHandler)
			advertisements.GET("/", advertisementHandler.GetAllAdvertisementsHandler)
			advertisements.PUT("/:id", advertisementHandler.UpdateAdvertisementHandler)
			advertisements.DELETE("/:id", advertisementHandler.DeleteAdvertisementHandler)
			advertisements.POST("/upload", advertisementHandler.UploadSheetAdvertisementHandler)
		}

		// Authentication routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authenticationHandler.RegisterHandler)
			auth.POST("/login", authenticationHandler.LoginHandler)
			auth.GET("/protected", authenticationHandler.ProtectedHandler)
		}
	}
}
