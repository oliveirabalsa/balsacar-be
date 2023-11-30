package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"github.com/oliveirabalsa/balsacar-be/internal/service"
)

type AuthenticationHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthenticationHandler {
	return &AuthenticationHandler{authService: authService}
}

func (h *AuthenticationHandler) RegisterHandler(c *gin.Context) {
	user := &entity.User{}
	err := h.authService.Register(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User registration failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthenticationHandler) LoginHandler(c *gin.Context) {
	user := &entity.User{}
	token, err := h.authService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func (h *AuthenticationHandler) ProtectedHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Protected route handler"})
}
