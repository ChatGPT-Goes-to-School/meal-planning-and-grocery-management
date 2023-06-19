package repositories

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/gorm"
)

// ListItemRepository represents the repository for the ListItem entity.
type ListItemRepository struct {
	db *gorm.DB
}

// NewListItemRepository creates a new instance of ListItemRepository.
func NewListItemRepository(db *gorm.DB) *ListItemRepository {
	return &ListItemRepository{
		db: db,
	}
}

// CreateListItem creates a new list item record in the database.
func (r *ListItemRepository) CreateListItem(listItem *models.ListItem) error {
	return r.db.Create(listItem).Error
}

// GetListItemByID retrieves a list item record from the database by its ID.
func (r *ListItemRepository) GetListItemByID(listItemID int) (*models.ListItem, error) {
	listItem := new(models.ListItem)
	if err := r.db.First(listItem, listItemID).Error; err != nil {
		return nil, err
	}
	return listItem, nil
}

// UpdateListItem updates an existing list item record in the database.
func (r *ListItemRepository) UpdateListItem(listItem *models.ListItem) error {
	return r.db.Save(listItem).Error
}

// DeleteListItem deletes a list item record from the database by its ID.
func (r *ListItemRepository) DeleteListItem(listItemID int) error {
	return r.db.Delete(&models.ListItem{}, listItemID).Error
}

func (r *ListItemRepository) GetListItemsByGroceryID(groceryID uint) ([]models.ListItem, error) {
	var listItems []models.ListItem
	res := r.db.Where("grocery_id = ?", groceryID).Find(&listItems)
	if res.Error != nil {
		return nil, res.Error
	}

	return listItems, nil
}
