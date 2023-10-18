package repository

import (
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
)

type AuthRepository interface {
	FindUserByID(userID uint) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type User struct {
	ID       uint
	Email    string
	Password string
}
