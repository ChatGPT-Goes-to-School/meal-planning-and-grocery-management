package handlers

import (
	"net/http"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/services"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GroceryHandler struct {
	service *services.GroceryService
	db      *gorm.DB
}

func NewGroceryHandler(db *gorm.DB) *GroceryHandler {
	service := services.NewGroceryService(db)
	return &GroceryHandler{
		service: service,
		db:      db,
	}
}

func (h *GroceryHandler) CreateGrocery(c *gin.Context) {
	var grocery models.Grocery
	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdGrocery, err := h.service.CreateGrocery(grocery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdGrocery)
}

func (h *GroceryHandler) GetGrocery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ConvertParamToInt(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	grocery, err := h.service.GetGroceryWithItems(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, grocery)
}

func (h *GroceryHandler) UpdateGrocery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ConvertParamToInt(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var grocery models.Grocery
	if err := c.ShouldBindJSON(&grocery); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedGrocery, err := h.service.UpdateGrocery(id, grocery)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedGrocery)
}

func (h *GroceryHandler) DeleteGrocery(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ConvertParamToInt(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.service.DeleteGrocery(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Grocery deleted successfully"})
}

func (h *GroceryHandler) GetGroceriesByUsername(c *gin.Context) {
	username := c.Param("username")

	// Call the service to retrieve the groceries by username
	groceries, err := h.service.GetGroceriesByUsername(username)
	if err != nil {
		// Handle the error
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the groceries as a JSON response
	c.JSON(http.StatusOK, groceries)
}
