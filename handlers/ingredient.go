package handlers

import (
	"net/http"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/services"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IngredientHandler struct {
	service *services.IngredientService
}

func NewIngredientService(db *gorm.DB) *GroceryHandler {
	service := services.NewIngredientService(db)
	return &GroceryHandler{
		service: service,
	}
}

func (h *GroceryHandler) AddIngredientToGrocery(c *gin.Context) {
	// Get grocery ID from request parameters
	groceryID, err := utils.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse ingredient data  from request body
	var ingredient []models.Ingredient
	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the service method to add the ingredient to the grocery
	err2 := h.service.AddIngredientToGrocery(groceryID, ingredient)
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(http.StatusCreated, nil)
}
