package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Variant struct {
	ID          int    `gorm:"primaryKey"`
	VariantName string `gorm:"not null;unique"`
	Quantity    int    `gorm:"not null"`
	ProductID   int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (v *Variant) BeforeCreate(tx *gorm.DB) (err error) {
	if v.Quantity <= 0 {
		err = errors.New("can't save product with quantity minimum is 1")
	}
	return
}
