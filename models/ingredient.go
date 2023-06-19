package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name        string  `gorm:"column:name"`
	Description string  `gorm:"column:description"`
	Category    string  `gorm:"column:category"`
	Quantity    int     `gorm:"column:quantity"`
	Price       float64 `gorm:"column:price"`
}
