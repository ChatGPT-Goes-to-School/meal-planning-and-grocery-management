package repositories

import (
	"fmt"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/config"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/gorm"
)

// MealPlanRepository handles the data storage and retrieval for meal plans
type MealRepository struct {
	// Any database or external service connections can be added here
		db *gorm.DB
}

func NewMealRepository() *MealRepository {
	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return nil
	}
	
	return &MealRepository{
		db: db,
	}
}

func (r *MealRepository) CreateMeal(db *gorm.DB, meal *models.Meal) error {
    return db.Create(meal).Error
}

func (r *MealRepository) GetMeal(db *gorm.DB, id uint) (*models.Meal, error) {
    var meal models.Meal
    err := db.First(&meal, id).Error
    if err != nil {
        return nil, err
    }
    return &meal, nil
}

func (r *MealRepository) UpdateMeal(db *gorm.DB, meal *models.Meal) error {
    return db.Save(meal).Error
}

func (r *MealRepository) DeleteMeal(db *gorm.DB, id uint) error {
    return db.Delete(&models.Meal{}, id).Error
}