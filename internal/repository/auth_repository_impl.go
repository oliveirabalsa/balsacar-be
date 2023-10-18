package repository

import (
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"gorm.io/gorm"
)

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{db: db}
}

func (ar *AuthRepositoryImpl) FindUserByID(userID uint) (*entity.User, error) {
	var user entity.User
	if err := ar.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *AuthRepositoryImpl) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := ar.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ar *AuthRepositoryImpl) CreateUser(user *entity.User) error {
	if err := ar.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
