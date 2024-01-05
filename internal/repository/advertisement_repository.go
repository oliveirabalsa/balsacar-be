package repository

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/balsacar-be/internal/dto"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
)

type AdvertisementRepository interface {
	Save(advertisement *entity.Advertisement) *entity.Advertisement
	Update(advertisementId uuid.UUID, advertisement *entity.Advertisement) (*entity.Advertisement, error)
	Delete(advertisementId uuid.UUID) error
	FindById(advertisementId uuid.UUID) (*entity.Advertisement, error)
	FindAll(filters *dto.AdvertisementParamsDto) []*entity.Advertisement
}
