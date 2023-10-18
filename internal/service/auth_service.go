package service

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}
