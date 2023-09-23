package entity

import (
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
	CreatedAt    time.Time        `gorm:"created_at" json:"createdAt"`
	UpdatedAt    time.Time        `gorm:"updated_at" json:"updatedAt"`
}

func (Advertisement) TableName() string {
	return "advertisements"
}
