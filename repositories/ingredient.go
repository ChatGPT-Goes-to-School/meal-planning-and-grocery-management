package repositories

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/gorm"
)

// IngredientRepository represents the repository for the Ingredient entity.
type IngredientRepository struct {
	db *gorm.DB
}

// NewIngredientRepository creates a new instance of IngredientRepository.
func NewIngredientRepository(db *gorm.DB) *IngredientRepository {
	return &IngredientRepository{
		db: db,
	}
}

// CreateIngredient creates a new ingredient record in the database.
func (r *IngredientRepository) CreateIngredient(ingredient *models.Ingredient) (uint, error) {
	item := r.db.Create(ingredient)
	return ingredient.ID, item.Error
}

// GetIngredientByID retrieves an ingredient record from the database by its ID.
func (r *IngredientRepository) GetIngredientByID(ingredientID int) (*models.Ingredient, error) {
	ingredient := new(models.Ingredient)
	if err := r.db.First(ingredient, ingredientID).Error; err != nil {
		return nil, err
	}
	return ingredient, nil
}

// UpdateIngredient updates an existing ingredient record in the database.
func (r *IngredientRepository) UpdateIngredient(ingredient *models.Ingredient) error {
	return r.db.Save(ingredient).Error
}

// DeleteIngredient deletes an ingredient record from the database by its ID.
func (r *IngredientRepository) DeleteIngredient(ingredientID int) error {
	return r.db.Delete(&models.Ingredient{}, ingredientID).Error
}
