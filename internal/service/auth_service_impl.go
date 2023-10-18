package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"github.com/oliveirabalsa/balsacar-be/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	secretKey      []byte
}

func NewAuthService(AuthRepository repository.AuthRepository, secretKey []byte) AuthService {
	return &AuthServiceImpl{
		AuthRepository: AuthRepository,
		secretKey:      secretKey,
	}
}

func (as *AuthServiceImpl) Register(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entity.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := as.AuthRepository.CreateUser(user); err != nil {
		return err
	}

	return nil
}

func (as *AuthServiceImpl) Login(email, password string) (string, error) {
	user, err := as.AuthRepository.FindUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["userID"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time

	tokenString, err := token.SignedString(as.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *AuthServiceImpl) VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return as.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
