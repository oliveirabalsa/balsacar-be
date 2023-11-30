package router

import (
	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/balsacar-be/internal/handler"
)

func InitRouter(router *gin.Engine, advertisementHandler *handler.AdvertisementHandler, authenticationHandler *handler.AuthenticationHandler) {
	api := router.Group("/api")
	{
		// Advertisement routes
		api.POST("/advertisements", advertisementHandler.CreateAdvertisementHandler)
		api.GET("/advertisements/:id", advertisementHandler.GetAdvertisementByIDHandler)
		api.GET("/advertisements", advertisementHandler.GetAllAdvertisementsHandler)
		api.PUT("/advertisements/:id", advertisementHandler.UpdateAdvertisementHandler)
		api.DELETE("/advertisements/:id", advertisementHandler.DeleteAdvertisementHandler)
		api.POST("/advertisements/upload", advertisementHandler.UploadSheetAdvertisementHandler)

		// Authentication routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", authenticationHandler.RegisterHandler)
			auth.POST("/login", authenticationHandler.LoginHandler)
			auth.GET("/protected", authenticationHandler.ProtectedHandler)
		}
	}
}
