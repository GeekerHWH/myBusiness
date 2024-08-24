package database

import "gorm.io/gorm"

type Goods struct {
	gorm.Model
	Name         string  `gorm:"not null"`
	NewPrice     float32 `gorm:"not null"`
	DiscountRate float32 `gorm:"not null"`
}
