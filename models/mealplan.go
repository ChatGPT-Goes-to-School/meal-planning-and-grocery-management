package models

import (
	"gorm.io/gorm"
)

// MealPlan represents a meal plan structure
type MealPlan struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Meals       string `json:"meals"`
	Username    string `json:"username"`
}
