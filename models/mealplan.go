package models

import (
	"time"

	"gorm.io/gorm"
)

// MealPlan represents a meal plan structure
type MealPlan struct {
	gorm.Model
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Duration    int        `json:"duration"`
	Meals       string     `json:"meals"`
	Username    string     `json:"username"`
	CreatedAt   *time.Time `gorm:"column:created_at;type:datetime"` // Specify the column type
	UpdatedAt   *time.Time `gorm:"column:updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at;index"` // Give a index name `deleted_at`
}
