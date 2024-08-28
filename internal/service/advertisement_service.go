package service

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/balsacar-be/internal/dto"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
)

type AdvertisementService interface {
	CreateAdvertisement(advertisement *entity.Advertisement) (*entity.Advertisement, error)
	UpdateAdvertisement(advertisementID uuid.UUID, advertisement *entity.Advertisement) (*entity.Advertisement, error)
	GetAdvertisementByID(advertisementID uuid.UUID) (*entity.Advertisement, error)
	GetAllAdvertisements(filters *dto.AdvertisementParamsDto) ([]*entity.Advertisement, int64, error)
	DeleteAdvertisement(advertisementID uuid.UUID) error
}
