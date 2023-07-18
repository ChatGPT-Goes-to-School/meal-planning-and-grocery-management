package models

import (
	"time"

	"gorm.io/gorm"
)

// MealPlan represents a meal plan structure
type MealPlan struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Breakfast   string    `json:"breakfast"`
	Lunch       string    `json:"lunch"`
	Dinner      string    `json:"dinner"`
	Snack       string    `json:"snack"`
	Username    string    `json:"username"`
	DatePlan    time.Time `json:"date"`
}
