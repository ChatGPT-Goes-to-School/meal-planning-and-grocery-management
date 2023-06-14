package models

// MealPlan represents a meal plan structure
type MealPlan struct {
	ID       int   `json:"id"`
	Name     string   `json:"name"`
	Description string `json:"description"`
	Duration int      `json:"duration"`
	Meals    []Meal `json:"meals"`
}