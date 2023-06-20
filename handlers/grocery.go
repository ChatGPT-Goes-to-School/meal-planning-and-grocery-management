package handlers

import (
	"net/http"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/services"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
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

// CreateGrocery creates a new grocery item.
// @Summary Create a grocery item
// @Description Create a new grocery item
// @Tags groceries
// @Accept json
// @Produce json
// @Param grocery body models.Grocery true "Grocery item object"
// @Success 200 {object} models.Grocery
// @Failure 400 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /groceries [post]
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

// GetGrocery retrieves a grocery item by ID.
// @Summary Get a grocery item
// @Description Retrieve a grocery item by ID
// @Tags groceries
// @Accept json
// @Produce json
// @Param id path int true "Grocery item ID"
// @Success 200 {object} models.Grocery
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /groceries/{id} [get]
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

// UpdateGrocery updates a grocery item by ID.
// @Summary Update a grocery item
// @Description Update a grocery item by ID
// @Tags groceries
// @Accept json
// @Produce json
// @Param id path int true "Grocery item ID"
// @Param grocery body models.Grocery true "Grocery item object"
// @Success 200 {object} models.Grocery
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /groceries/{id} [put]
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

// DeleteGrocery deletes a grocery item by ID.
// @Summary Delete a grocery item
// @Description Delete a grocery item by ID
// @Tags groceries
// @Param id path int true "Grocery item ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /groceries/{id} [delete]
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

// GetGroceriesByUsername retrieves the groceries associated with a specific username.
// @Summary Retrieve groceries by username
// @Description Retrieves the groceries associated with a specific username
// @Tags groceries
// @Param username path string true "Username"
// @Success 200 {array} models.Grocery
// @Failure 500 {object} httputil.HTTPError
// @Router /groceries/username/{username} [get]
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
