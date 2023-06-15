package repositories

import (
	"fmt"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/config"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/errors"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/gorm"
)

// MealPlanRepository handles the data storage and retrieval for meal plans
type MealPlanRepository struct {
	// Any database or external service connections can be added here
		db *gorm.DB
}

func NewMealPlanRepository() *MealPlanRepository {
	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil
	}
	
	return &MealPlanRepository{
		db: db,
	}
}

// CreateMealPlan creates a new meal plan in the repository
func (r *MealPlanRepository) CreateMealPlan(mealPlan models.MealPlan) (models.MealPlan, error) {
    err := r.db.Create(&mealPlan).Error
    if err != nil {
        return models.MealPlan{}, err
    }

    return mealPlan, nil
}

// GetMealPlan retrieves a meal plan from the repository by ID
func (r *MealPlanRepository) GetMealPlan(id int) (models.MealPlan, error) {
	// Implementation to fetch the meal plan from the database
	var mealPlan models.MealPlan
    err := r.db.First(&mealPlan, id).Error
    if err != nil {
        if (err == gorm.ErrRecordNotFound) {
            return models.MealPlan{}, errors.NewMealPlanNotFoundError(id)
        }
        return models.MealPlan{}, err
    }
    return mealPlan, nil
}

// UpdateMealPlan updates a meal plan in the repository by ID
func (r *MealPlanRepository) UpdateMealPlan(id int, mealPlan models.MealPlan) error {
	// Implementation to update the meal plan in the database
	result := r.db.Save(&mealPlan)
    if result.Error != nil {
        return errors.NewMealPlanNotFoundError(id)
    }
    return nil
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

