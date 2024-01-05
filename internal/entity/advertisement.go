package entity

import (
	"strconv"
	"strings"
	"time"

	"github.com/lib/pq"
)

// TransmissionEnum represents the possible transmission types.
type TransmissionEnum string

const (
	Manual    TransmissionEnum = "manual"
	Automatic TransmissionEnum = "automatic"
)

type CarTypeEnum string

const (
	SUV         CarTypeEnum = "suv"
	Sedan       CarTypeEnum = "sedan"
	Hatchback   CarTypeEnum = "hatchback"
	Coupe       CarTypeEnum = "coupe"
	Convertible CarTypeEnum = "convertible"
)

type FuelTypeEnum string

const (
	Petrol   FuelTypeEnum = "petrol"
	Diesel   FuelTypeEnum = "diesel"
	Hybrid   FuelTypeEnum = "hybrid"
	Electric FuelTypeEnum = "electric"
)

type StatusEnum string

const (
	Sold      StatusEnum = "sold"
	Available StatusEnum = "available"
	Reserved  StatusEnum = "reserved"
)

type Advertisement struct {
	ID           string           `gorm:"primaryKey" json:"id"`
	Transmission TransmissionEnum `json:"transmission"`
	Type         CarTypeEnum      `json:"type"`
	City         string           `json:"city"`
	Price        float64          `json:"price"`
	Make         string           `json:"make"`
	Model        string           `json:"model"`
	Year         int              `json:"year"`
	Kilometerage int              `json:"kilometerage"`
	Fuel         FuelTypeEnum     `json:"fuel"`
	Description  string           `json:"description"`
	Contact      string           `json:"contact"`
	Images       pq.StringArray   `gorm:"type:text[]" json:"images"`
	Status       StatusEnum       `json:"status"`
	Active       bool             `gorm:"default: false" json:"active"`
	BestOffer    bool             `gorm:"default: false" json:"best_offer"`
	CreatedAt    time.Time        `gorm:"created_at" json:"createdAt"`
	UpdatedAt    time.Time        `gorm:"updated_at" json:"updatedAt"`
}

func (Advertisement) TableName() string {
	return "advertisements"
}

func (*Advertisement) FromKeyValue(ad *Advertisement, key, value string) {
	switch key {
	case "Transmission":
		ad.Transmission = TransmissionEnum(value)
	case "Type":
		ad.Type = CarTypeEnum(value)
	case "City":
		ad.City = value
	case "Price":
		price, err := strconv.ParseFloat(value, 10)
		if err == nil {
			ad.Price = price
		}
	case "Make":
		ad.Make = value
	case "Model":
		ad.Model = value
	case "Year":
		year, err := strconv.Atoi(value)
		if err == nil {
			ad.Year = year
		}
	case "Kilometerage":
		kilometerage, err := strconv.Atoi(value)
		if err == nil {
			ad.Kilometerage = kilometerage
		}
	case "Fuel":
		ad.Fuel = FuelTypeEnum(value)
	case "Description":
		ad.Description = value
	case "Contact":
		ad.Contact = value
	case "Images":
		ad.Images = pq.StringArray(strings.Split(value, ","))
	case "Status":
		ad.Status = StatusEnum(value)
	case "Active":
		parsedValue, _ := strconv.ParseBool(value)
		ad.Active = parsedValue
	case "BestOffer":
		parsedValue, _ := strconv.ParseBool(value)
		ad.BestOffer = parsedValue
	}
}
