package repository

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
)

type AuthRepository interface {
	FindUserByID(userID *uuid.UUID) (*entity.User, error)
	FindUserByEmail(email string) (*entity.User, error)
	CreateUser(user *entity.User) error
}

type User struct {
	ID       *uuid.UUID
	Email    string
	Password string
}
