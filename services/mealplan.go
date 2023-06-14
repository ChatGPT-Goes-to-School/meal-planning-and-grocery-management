package services

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/errors"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
)

var mealPlans []models.MealPlan

// MealPlanService handles the business logic for meal plans
type MealPlanService struct {
	// Any dependencies or repositories needed by the service can be added here
}

// Create a meal plan
func (s *MealPlanService) CreateMealPlan(mealPlan models.MealPlan) models.MealPlan {
	// Assign a unique ID to the meal plan
	mealPlans = append(mealPlans, mealPlan)

	return mealPlan
}

// Get a meal plan by ID
func (s *MealPlanService) GetMealPlan(id int) (models.MealPlan, error) {
    for _, mp := range mealPlans {
        if mp.ID == id {
            return mp, nil
        }
    }

    return models.MealPlan{}, errors.NewMealPlanNotFoundError(id)
}

// Update a meal plan by ID
func (s *MealPlanService) UpdateMealPlan(id int, updatedMealPlan models.MealPlan) (models.MealPlan, error) {
	for i, mp := range mealPlans {
		if mp.ID == id {
			// Update the meal plan
			mealPlans[i] = updatedMealPlan
			return updatedMealPlan, nil
		}
	}

	return models.MealPlan{}, errors.NewMealPlanNotFoundError(id)
}

// Delete a meal plan by ID
func (s *MealPlanService) DeleteMealPlan(id int) error {
	for i, mp := range mealPlans {
		if mp.ID == id {
			// Remove the meal plan
			mealPlans = append(mealPlans[:i], mealPlans[i+1:]...)
			return nil
		}
	}

	return errors.NewMealPlanNotFoundError(id)
}