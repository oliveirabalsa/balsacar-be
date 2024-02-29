package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/balsacar-be/docs"
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
	"github.com/oliveirabalsa/balsacar-be/internal/middleware"
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

	router.Use(middleware.CORSMiddleware())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api")
	{
		advertisements := api.Group("/advertisements")
		{
			advertisements.GET("/all", advertisementHandler.GetAllAdvertisementsHandler)
			advertisements.GET("/:id", advertisementHandler.GetAdvertisementByIDHandler)
			advertisements.POST("/", authMiddleware, advertisementHandler.CreateAdvertisementHandler)
			advertisements.PUT("/:id", authMiddleware, advertisementHandler.UpdateAdvertisementHandler)
			advertisements.DELETE("/:id", authMiddleware, advertisementHandler.DeleteAdvertisementHandler)
			advertisements.POST("/upload", authMiddleware, advertisementHandler.UploadSheetAdvertisementHandler)
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
