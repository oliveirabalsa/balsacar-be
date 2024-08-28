package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
)

type AdvertisementParamsDto struct {
	YearFrom     string                  `form:"year_from"`
	YearTo       string                  `form:"year_to" `
	Model        string                  `form:"model" validate:"required,max=50"`
	City         string                  `form:"city"`
	Type         entity.CarTypeEnum      `form:"type"`
	Transmission entity.TransmissionEnum `form:"transmission"`
	BestOffer    bool                    `form:"best_offer"`
	Highlight    bool                    `form:"highlight"`
	Page         int64                   `form:"page"`
	PageSize     int64                   `form:"page_size"`
}

type Paginated struct {
	TotalPages int64       `json:"total_pages"`
	Total      int64       `json:"total"`
	Page       int64       `json:"page"`
	PageSize   int64       `json:"page_size"`
	Data       interface{} `json:"data"`
}

func (apd AdvertisementParamsDto) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("currentYear", func(fl validator.FieldLevel) bool {
		currentYear := time.Now().Year()
		return int(fl.Field().Int()) <= currentYear
	})

	return validate.Struct(apd)
}
