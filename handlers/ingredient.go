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

func NewIngredientHandler(db *gorm.DB) *IngredientHandler {
	service := services.NewIngredientService(db)
	return &IngredientHandler{
		service: service,
	}
}

// AddIngredientToGrocery adds an ingredient to a grocery.
// @Summary Add ingredient to grocery
// @Description Add ingredient to a grocery by ID
// @Tags Grocery
// @Accept json
// @Produce json
// @Param id path int true "Grocery ID"
// @Param ingredient body []models.Ingredient true "Ingredient object"
// @Success 201 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grocery/{id}/ingredient [post]
func (h *IngredientHandler) AddIngredientToGrocery(c *gin.Context) {
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

// RemoveIngredientFromGrocery removes an ingredient from a grocery.
// @Summary Remove ingredient from grocery
// @Description Remove ingredient from a grocery by grocery ID and ingredient ID
// @Tags Grocery
// @Accept json
// @Produce json
// @Param groceryID path int true "Grocery ID"
// @Param ingredientID path int true "Ingredient ID"
// @Success 200 {object} string "OK"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grocery/{groceryID}/ingredient/{ingredientID} [delete]
func (h *IngredientHandler) RemoveIngredientFromGrocery(c *gin.Context) {
	groceryID, err := utils.ConvertParamToInt(c.Param("groceryID"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ingredientID, err2 := utils.ConvertParamToInt(c.Param("ingredientID"))
	if err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	err3 := h.service.RemoveIngredientFromGrocery(groceryID, ingredientID)
	if err3 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err3.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ingredient removed successfully"})
}

// UpdateIngredientQuantity updates the quantity of an ingredient.
// @Summary Update ingredient quantity
// @Description Update the quantity of an ingredient
// @Tags Grocery
// @Accept json
// @Produce json
// @Param ingredient body models.Ingredient true "Ingredient object"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /grocery/ingredient [put]
func (h *IngredientHandler) UpdateIngredientQuantity(c *gin.Context) {
	var ingredient models.Ingredient

	if err := c.ShouldBindJSON(&ingredient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateIngredientQuantity(ingredient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ingredient quantity updated successfully"})
}
