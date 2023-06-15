package handlers

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/repositories"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/services"
)

// MealPlanHandler handles the HTTP requests for meal plans
type MealHandler struct {
	service    services.MealPlanService
	repository repositories.MealRepository
}
