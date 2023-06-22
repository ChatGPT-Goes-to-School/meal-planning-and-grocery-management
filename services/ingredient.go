package services

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/repositories"
	"gorm.io/gorm"
)

type IngredientService struct {
	ingredientRepository *repositories.IngredientRepository
	listItemRepository   *repositories.ListItemRepository
	groceryRepository    *repositories.GroceryRepository
}

func NewIngredientService(db *gorm.DB) *IngredientService {
	ingredientRepository := repositories.NewIngredientRepository(db)
	listItemRepository := repositories.NewListItemRepository(db)
	groceryRepository := repositories.NewGroceryRepository(db)

	return &IngredientService{
		ingredientRepository: ingredientRepository,
		listItemRepository:   listItemRepository,
		groceryRepository:    groceryRepository,
	}
}

func (s *IngredientService) AddIngredientToGrocery(groceryID int, ingredient []models.Ingredient) error {
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

func (s *IngredientService) RemoveIngredientFromGrocery(groceryID, ingredientID int) error {
	res, err := s.listItemRepository.GetListItemByGroceryIDAndIngredientID(uint(groceryID), uint(ingredientID))
	if err != nil {
		return err
	}

	err2 := s.listItemRepository.DeleteListItem(int(res.ID))
	if err2 != nil {
		// Handle any repository errors
		return err2
	}

	err3 := s.ingredientRepository.DeleteIngredient(ingredientID)
	if err3 != nil {
		// Handle any repository errors
		return err3
	}

	return nil
}

func (s *IngredientService) UpdateIngredientQuantity(ingredient models.Ingredient) error {
	// Convert groceryID and ingredientID to appropriate data types if needed

	err := s.ingredientRepository.UpdateIngredient(&ingredient)
	if err != nil {
		// Handle any repository errors
		return err
	}

	return nil
}
