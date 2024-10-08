package service

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/balsacar-be/internal/dto"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"github.com/oliveirabalsa/balsacar-be/internal/repository"
)

type AdvertisementServiceImpl struct {
	advertisementRepository repository.AdvertisementRepository
}

func NewAdvertisementService(advertisementRepository repository.AdvertisementRepository) AdvertisementService {
	return &AdvertisementServiceImpl{
		advertisementRepository: advertisementRepository,
	}
}

func (s *AdvertisementServiceImpl) CreateAdvertisement(advertisement *entity.Advertisement) (*entity.Advertisement, error) {
	advertisement.ID = uuid.New().String()
	return s.advertisementRepository.Save(advertisement), nil
}

func (s *AdvertisementServiceImpl) UpdateAdvertisement(advertisementID uuid.UUID, advertisement *entity.Advertisement) (*entity.Advertisement, error) {
	updatedAdvertisement, err := s.advertisementRepository.Update(advertisementID, advertisement)
	if err != nil {
		return nil, err
	}
	return updatedAdvertisement, nil
}

func (s *AdvertisementServiceImpl) GetAdvertisementByID(advertisementID uuid.UUID) (*entity.Advertisement, error) {
	return s.advertisementRepository.FindById(advertisementID)
}

func (s *AdvertisementServiceImpl) GetAllAdvertisements(filters *dto.AdvertisementParamsDto) ([]*entity.Advertisement, int64, error) {
	return s.advertisementRepository.FindAll(filters)
}

func (s *AdvertisementServiceImpl) DeleteAdvertisement(advertisementID uuid.UUID) error {
	return s.advertisementRepository.Delete(advertisementID)
}
