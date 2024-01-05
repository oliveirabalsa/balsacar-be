package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/oliveirabalsa/balsacar-be/internal/entity"
)

type AdvertisementParamsDto struct {
	YearFrom     string                  `form:"year_from" `
	YearTo       string                  `form:"year_to" `
	Model        string                  `form:"model"`
	City         string                  `form:"city"`
	Type         entity.CarTypeEnum      `form:"type"`
	Transmission entity.TransmissionEnum `form:"transmission"`
}

func (apd AdvertisementParamsDto) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("currentYear", func(fl validator.FieldLevel) bool {
		currentYear := time.Now().Year()
		return int(fl.Field().Int()) <= currentYear
	})

	return validate.Struct(apd)
}
