package repository

import (
	"github.com/google/uuid"
	"github.com/oliveirabalsa/balsacar-be/internal/dto"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
	"gorm.io/gorm"
)

type AdvertisementRepositoryImpl struct {
	db *gorm.DB
}

func NewAdvertisementRepository(db *gorm.DB) AdvertisementRepository {
	return &AdvertisementRepositoryImpl{
		db: db,
	}
}

func (r *AdvertisementRepositoryImpl) Delete(advertisementId uuid.UUID) error {
	return r.db.Delete(&entity.Advertisement{}, advertisementId).Error
}

func (r *AdvertisementRepositoryImpl) FindAll(filters *dto.AdvertisementParamsDto) []*entity.Advertisement {
	var advertisements []*entity.Advertisement
	query := r.db.Model(&entity.Advertisement{})

	generateFilteredGetQuery(query, filters)

	query.Find(&advertisements)

	return advertisements
}

func (r *AdvertisementRepositoryImpl) FindById(advertisementId uuid.UUID) (*entity.Advertisement, error) {
	var advertisement *entity.Advertisement
	err := r.db.First(&advertisement, advertisementId).Error
	return advertisement, err
}

func (r *AdvertisementRepositoryImpl) Save(advertisement *entity.Advertisement) *entity.Advertisement {
	r.db.Create(&advertisement)
	return advertisement
}

func (r *AdvertisementRepositoryImpl) Update(advertisementId uuid.UUID, updates *entity.Advertisement) (*entity.Advertisement, error) {
	advertisement := &entity.Advertisement{}

	// Find the advertisement by ID
	err := r.db.First(&advertisement, "id = ?", advertisementId).Error
	if err != nil {
		return nil, err
	}

	err = r.db.Model(&advertisement).Updates(updates).Error
	if err != nil {
		return nil, err
	}

	return advertisement, nil
}

func generateFilteredGetQuery(query *gorm.DB, filters *dto.AdvertisementParamsDto) *gorm.DB {
	if filters.YearFrom != "" {
		query = query.Where("year >= ?", filters.YearFrom)
	}
	if filters.YearTo != "" {
		query = query.Where("year <= ?", filters.YearTo)
	}
	if filters.Model != "" {
		query = query.Where("LOWER(model) LIKE LOWER(?)", "%"+filters.Model+"%")
	}
	if filters.Type != "" {
		query = query.Where("LOWER(type) LIKE LOWER(?)", filters.Type)
	}
	if filters.City != "" {
		query = query.Where("LOWER(city) LIKE LOWER(?)", "%"+filters.City+"%")
	}
	if filters.Transmission != "" {
		query = query.Where("LOWER(transmission) LIKE LOWER(?)", "%"+filters.Transmission+"%")
	}
	if filters.BestOffer {
		query = query.Where("best_offer = ?", true)
	}

	return query
}
