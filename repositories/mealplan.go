package repositories

import (
	"fmt"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/errors"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/gorm"
)

// MealPlanRepository handles the data storage and retrieval for meal plans
type MealPlanRepository struct {
	// Any database or external service connections can be added here
	db *gorm.DB
}

func NewMealPlanRepository(db *gorm.DB) *MealPlanRepository {
	return &MealPlanRepository{
		db: db,
	}
}

// CreateMealPlan creates a new meal plan in the repository
func (r *MealPlanRepository) CreateMealPlan(mealPlan models.MealPlan) (models.MealPlan, error) {
	fmt.Println("Creating meal plan", mealPlan)
	fmt.Println("db", r.db)
	err := r.db.Create(&mealPlan).Error
	if err != nil {
		fmt.Println("Failed to create meal plan:", err)
		return models.MealPlan{}, err
	}

	return mealPlan, nil
}

// GetMealPlan retrieves a meal plan from the repository by Username
func (r *MealPlanRepository) GetMealPlanByUsername(username string) ([]models.MealPlan, error) {
	// Implementation to fetch the meal plan from the database
	var mealPlan []models.MealPlan
	res := r.db.Where("username = ?", username).Find(&mealPlan)
	if res.Error != nil {
		return nil, errors.NewMealPlanNotFoundError(0)
	}

	return mealPlan, nil
}

// GetMealPlan retrieves a meal plan from the repository by ID
func (r *MealPlanRepository) GetMealPlan(id int) (models.MealPlan, error) {
	// Implementation to fetch the meal plan from the database
	var mealPlan models.MealPlan
	res := r.db.First(&mealPlan, id)

	if res.Error != nil {
		return models.MealPlan{}, errors.NewMealPlanNotFoundError(id)
	}

	return mealPlan, res.Error
}

// UpdateMealPlan updates a meal plan in the repository by ID
func (r *MealPlanRepository) UpdateMealPlan(id int, mealPlan models.MealPlan) (models.MealPlan, error) {
	// Implementation to update the meal plan in the database
	result := r.db.Save(&mealPlan)
	if result.Error != nil {
		return models.MealPlan{}, errors.NewMealPlanNotFoundError(id)
	}
	return mealPlan, nil
}

// DeleteMealPlan deletes a meal plan from the repository by ID
func (r *MealPlanRepository) DeleteMealPlan(id int) error {
	mealPlan := models.MealPlan{}
	result := r.db.Where("id = ?", id).Delete(&mealPlan)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
