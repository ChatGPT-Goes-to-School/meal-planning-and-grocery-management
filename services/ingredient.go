package services

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/repositories"
	"gorm.io/gorm"
)

type IngredientService struct {
	ingredientRepository repositories.IngredientRepository
	listItemRepository   repositories.ListItemRepository
	groceryRepository    repositories.GroceryRepository
}

func NewIngredientService(db *gorm.DB) *GroceryService {
	ingredientRepository := repositories.NewIngredientRepository(db)
	listItemRepository := repositories.NewListItemRepository(db)
	groceryRepository := repositories.NewGroceryRepository(db)

	return &GroceryService{
		ingredientRepository: ingredientRepository,
		listItemRepository:   listItemRepository,
		groceryRepository:    groceryRepository,
	}
}

func (s *GroceryService) AddIngredientToGrocery(groceryID int, ingredient []models.Ingredient) error {

	// Create the ingredient in the database
	for _, ingredient := range ingredient {
		id, err := s.ingredientRepository.CreateIngredient(&ingredient)
		if err != nil {
			return err
		}

		// Create the item listing
		listItem := models.ListItem{
			IngredientID: int(id),
			GroceryID:    groceryID,
		}

		err2 := s.listItemRepository.CreateListItem(&listItem)
		if err2 != nil {
			return err2
		}

		// Retrieve the grocery by ID
		grocery, err := s.groceryRepository.GetGroceryByID(groceryID)
		grocery.Items = append(grocery.Items, listItem)
		if err != nil {
			return err
		}
	}

	return nil
}
