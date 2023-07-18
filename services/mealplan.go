package services

import (
	"fmt"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/errors"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/repositories"
	"gorm.io/gorm"
)

type MealPlanService struct {
	repository *repositories.MealPlanRepository
	db         *gorm.DB
}

// NewMealPlanService creates a new instance of MealPlanService.
func NewMealPlanService(db *gorm.DB) *MealPlanService {
	repository := repositories.NewMealPlanRepository(db)
	fmt.Println("NewMealPlanService")
	return &MealPlanService{
		repository: repository,
		db:         db,
	}
}

// Create a meal plan
func (s *MealPlanService) CreateMealPlan(mealPlan models.MealPlan) (models.MealPlan, error) {
	res, err := s.repository.CreateMealPlan(mealPlan)
	if err != nil {
		return res, err
	}

	return res, err
}

// Get a meal plan by ID
func (s *MealPlanService) GetMealPlan(id int) (models.MealPlan, error) {
	res, err := s.repository.GetMealPlan(id)

	if err != nil {
		return res, errors.NewMealPlanNotFoundError(id)
	}

	return res, nil
}

// Get a meal plan by Username
func (s *MealPlanService) GetMealPlanByUsername(username string) ([]models.MealPlan, error) {
	res, err := s.repository.GetMealPlanByUsername(username)

	if err != nil {
		return res, errors.NewMealPlanNotFoundError(0)
	}

	return res, nil
}

// Update a meal plan by ID
func (s *MealPlanService) UpdateMealPlan(id int, updatedMealPlan models.MealPlan) (models.MealPlan, error) {
	mealPlan, err := s.GetMealPlan(id)

	if err != nil {
		return models.MealPlan{}, errors.NewMealPlanNotFoundError(id)
	}

	mealPlan.Description = updatedMealPlan.Description
	mealPlan.Name = updatedMealPlan.Name
	mealPlan.Duration = updatedMealPlan.Duration
	mealPlan.Breakfast = updatedMealPlan.Breakfast
	mealPlan.Lunch = updatedMealPlan.Lunch
	mealPlan.Dinner = updatedMealPlan.Dinner
	mealPlan.Snack = updatedMealPlan.Snack

	res, err := s.repository.UpdateMealPlan(id, mealPlan)

	if err != nil {
		return models.MealPlan{}, err
	}

	return res, nil
}

// Delete a meal plan by ID
func (s *MealPlanService) DeleteMealPlan(id int) error {
	err := s.repository.DeleteMealPlan(id)

	if err != nil {
		return errors.NewMealPlanNotFoundError(id)
	}

	return nil
}

func (s *MealPlanService) GetMealPlanByUsernameAndDatePlan(username string, datePlan string) (models.MealPlan, error) {
	res, err := s.repository.GetMealPlanByUsernameAndDatePlan(username, datePlan)

	if err != nil {
		return res, errors.NewMealPlanNotFoundError(0)
	}

	return res, nil
}
