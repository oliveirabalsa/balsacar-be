package repository

import (
	"fmt"

	"github.com/google/uuid"
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

func (r *AdvertisementRepositoryImpl) FindAll() []*entity.Advertisement {
	var advertisements []*entity.Advertisement
	r.db.Find(&advertisements)
	fmt.Println(advertisements)
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

func (r *AdvertisementRepositoryImpl) Update(advertisement *entity.Advertisement) *entity.Advertisement {
	r.db.Save(&advertisement)
	return advertisement
}
