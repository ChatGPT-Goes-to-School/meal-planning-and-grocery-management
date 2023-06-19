package services

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/repositories"
	"gorm.io/gorm"
)

// MealPlanService handles the business logic for meal plans
type GroceryService struct {
	// Any dependencies or repositories needed by the service can be added here
	groceryRepository    *repositories.GroceryRepository
	ingredientRepository *repositories.IngredientRepository
	listItemRepository   *repositories.ListItemRepository
	db                   *gorm.DB
}

// NewMealPlanService creates a new instance of MealPlanService.
func NewGroceryService(db *gorm.DB) *GroceryService {
	groceryRepository := repositories.NewGroceryRepository(db)
	ingredientRepository := repositories.NewIngredientRepository(db)
	listItemRepository := repositories.NewListItemRepository(db)
	return &GroceryService{
		groceryRepository:    groceryRepository,
		ingredientRepository: ingredientRepository,
		listItemRepository:   listItemRepository,
		db:                   db,
	}
}

func (s *GroceryService) GetGroceryWithItems(id int) (models.Grocery, error) {
	grocery, err := s.groceryRepository.GetGroceryByID(id)

	if err != nil {
		return models.Grocery{}, err
	}

	// Get the list items for the grocery
	listItems, err := s.listItemRepository.GetListItemsByGroceryID(grocery.ID)
	if err != nil {
		return models.Grocery{}, err
	}

	// Fetch the ingredient details for each list item
	for i, listItem := range listItems {
		ingredient, err := s.ingredientRepository.GetIngredientByID(listItem.IngredientID)
		if err != nil {
			return models.Grocery{}, err
		}
		listItems[i].Ingredient = *ingredient
	}

	// Assign the list items to the grocery
	grocery.Items = listItems

	return *grocery, nil
}

func (s *GroceryService) CreateGrocery(grocery models.Grocery) (models.Grocery, error) {
	// Save the grocery to the database
	createdGrocery, err := s.groceryRepository.CreateGrocery(&grocery)
	if err != nil {
		return models.Grocery{}, err
	}

	return *createdGrocery, nil
}

func (s *GroceryService) GetGroceryByID(id int) (models.Grocery, error) {
	grocery, err := s.groceryRepository.GetGroceryByID(id)
	if err != nil {
		return models.Grocery{}, err
	}
	return *grocery, nil
}

func (s *GroceryService) UpdateGrocery(id int, grocery models.Grocery) (*models.Grocery, error) {
	oldGrocery, err := s.GetGroceryByID(id)
	if err != nil {
		return nil, err
	}
	oldGrocery.Name = grocery.Name
	oldGrocery.Items = grocery.Items

	res, err := s.groceryRepository.UpdateGrocery(&oldGrocery)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *GroceryService) DeleteGrocery(id int) error {
	err := s.groceryRepository.DeleteGrocery(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *GroceryService) GetGroceriesByUsername(username string) ([]models.Grocery, error) {
	groceries, err := s.groceryRepository.GetGroceriesByUsername(username)
	if err != nil {
		return nil, err
	}
	return groceries, nil
}
