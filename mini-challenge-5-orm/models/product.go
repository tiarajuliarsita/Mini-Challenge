package models

import "time"

type Product struct {
	ID        int    `gorm:"primaryKey"`
	Name      string `gorm:"not nul;unique"`
	Variants  []Variant
	CreatedAt time.Time
	UpdatedAt time.Time
}
