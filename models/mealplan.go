package models

import "gorm.io/gorm"

// MealPlan represents a meal plan structure
type MealPlan struct {
	gorm.Model
	ID       int   `json:"id" gorm:"primaryKey"`
	Name     string   `json:"name"`
	Description string `json:"description"`
	Duration int      `json:"duration"`
	Meals    []int `json:"meals"`
}