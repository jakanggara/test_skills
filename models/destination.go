package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *DbDestination) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	s.CreatedAt, s.UpdatedAt = t, t
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}
func (s *DbDestination) BeforeUpdate(tx *gorm.DB) error {
	s.UpdatedAt = time.Now()
	return nil
}

type DbDestination struct {
	ID           uuid.UUID
	ProductName  string
	Qty          uint
	SellingPrice uint
	PromoPrice   uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
