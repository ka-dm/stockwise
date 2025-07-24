package models

import (
	"time"

	"gorm.io/gorm"
)

// Stock representa la estructura de un stock en la base de datos
type Stock struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	Ticker     string         `json:"ticker" gorm:"not null;uniqueIndex"`
	TargetFrom string         `json:"target_from"`
	TargetTo   string         `json:"target_to"`
	Company    string         `json:"company"`
	Action     string         `json:"action"`
	Brokerage  string         `json:"brokerage"`
	RatingFrom string         `json:"rating_from"`
	RatingTo   string         `json:"rating_to"`
	Time       string         `json:"time"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName especifica el nombre de la tabla
func (Stock) TableName() string {
	return "stocks"
}

// BeforeCreate hook para validaciones antes de crear
func (s *Stock) BeforeCreate(tx *gorm.DB) error {
	if s.Ticker == "" {
		return gorm.ErrInvalidData
	}
	return nil
}

// Validate valida los datos del stock
func (s *Stock) Validate() error {
	if s.Ticker == "" {
		return gorm.ErrInvalidData
	}
	return nil
}
