package models

import (
	"gorm.io/gorm"
)

type ListItem struct {
	gorm.Model
	IngredientID int        `gorm:"column:ingredient_id"`
	Ingredient   Ingredient `gorm:"foreignKey:IngredientID"`
	GroceryID    int        `gorm:"column:grocery_id"`
	Grocery      Grocery    `gorm:"foreignKey:GroceryID"`
	Quantity     int        `gorm:"column:quantity"`
}
