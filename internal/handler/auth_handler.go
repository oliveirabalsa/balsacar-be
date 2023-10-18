// handler/auth_handler.go
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
	user := &entity.User{} // Replace with the actual user data from the request
	err := h.authService.Register(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User registration failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthenticationHandler) LoginHandler(c *gin.Context) {
	// Implement user login logic
	// You can use the authService to authenticate the user and generate a token here
	user := &entity.User{} // Replace with the actual user data from the request
	token, err := h.authService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Login failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func (h *AuthenticationHandler) ProtectedHandler(c *gin.Context) {
	// Implement a protected route that requires authentication
	// You can use middleware to check the JWT token, and if it's valid, allow access
	// Otherwise, return an error response
	c.JSON(http.StatusOK, gin.H{"message": "Protected route handler"})
}
