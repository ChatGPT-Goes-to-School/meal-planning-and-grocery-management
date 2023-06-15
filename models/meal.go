package models

import (
	"gorm.io/gorm"
)


type Meal struct {
	gorm.Model
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"column:name"`
	Ingredients string  `gorm:"column:ingredients"`
	Instructions string   `gorm:"column:instructions"`
	MealPlanID  uint       `gorm:"column:meal_plan_id"` // Foreign key referencing MealPlan's primary key
}