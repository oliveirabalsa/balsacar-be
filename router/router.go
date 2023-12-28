package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
)

func InitRouter(router *gin.Engine, advertisementHandler *handler.AdvertisementHandler, authenticationHandler *handler.AuthenticationHandler, authMiddleware gin.HandlerFunc) {
	api := router.Group("/api")
	{
		// Advertisement routes
		advertisements := api.Group("/advertisements", authMiddleware)
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
