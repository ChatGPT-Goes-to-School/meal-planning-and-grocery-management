package repositories

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/gorm"
)

type GroceryRepository struct {
	// Any database or external service connections can be added here
	db *gorm.DB
}

func NewGroceryRepository(db *gorm.DB) *GroceryRepository {

	return &GroceryRepository{
		db: db,
	}
}

// CreateGrocery creates a new grocery record in the database.
func (r *GroceryRepository) CreateGrocery(grocery *models.Grocery) (*models.Grocery, error) {
	res := r.db.Create(grocery)
	if res.Error != nil {
		return nil, res.Error
	}

	return grocery, nil
}

// GetGroceryByID retrieves a grocery record from the database by its ID.
func (r *GroceryRepository) GetGroceryByID(groceryID int) (*models.Grocery, error) {
	grocery := new(models.Grocery)
	if err := r.db.First(grocery, groceryID).Error; err != nil {
		return nil, err
	}
	return grocery, nil
}

// UpdateGrocery updates an existing grocery record in the database.
func (r *GroceryRepository) UpdateGrocery(grocery *models.Grocery) (*models.Grocery, error) {
	res := r.db.Save(grocery)

	if res.Error != nil {
		return nil, res.Error
	}

	return grocery, nil
}

// DeleteGrocery deletes a grocery record from the database by its ID.
func (r *GroceryRepository) DeleteGrocery(groceryID int) error {
	return r.db.Delete(&models.Grocery{}, groceryID).Error
}

func (r *GroceryRepository) GetGroceriesByUsername(username string) ([]models.Grocery, error) {
	var groceries []models.Grocery
	if err := r.db.Where("username = ?", username).Find(&groceries).Error; err != nil {
		return nil, err
	}
	return groceries, nil
}
