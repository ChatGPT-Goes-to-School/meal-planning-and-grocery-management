package handlers

import (
	"net/http"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/services"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// MealPlanHandler handles the HTTP requests for meal plans
type MealPlanHandler struct {
	service *services.MealPlanService
	db      *gorm.DB
}

// NewMealPlanHandler creates a new instance of MealPlanHandler
func NewMealPlanHandler(db *gorm.DB) *MealPlanHandler {
	service := services.NewMealPlanService(db)
	return &MealPlanHandler{
		service: service,
		db:      db,
	}
}

// CreateMealPlan handles the creation of a meal plan
func (h *MealPlanHandler) CreateMealPlan(c *gin.Context) {
	var mealPlan models.MealPlan
	if err := c.ShouldBindJSON(&mealPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMealPlan, err := h.service.CreateMealPlan(mealPlan)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdMealPlan)
}

// GetMealPlan handles the retrieval of a meal plan by ID
func (h *MealPlanHandler) GetMealPlan(c *gin.Context) {
	id, err := utils.ConvertParamToInt(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mealPlan, err := h.service.GetMealPlan(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mealPlan)
}

// GetMealPlan handles the retrieval of a meal plan by ID
func (h *MealPlanHandler) GetMealPlanByUsername(c *gin.Context) {
	username := c.Param("username")

	mealPlan, err := h.service.GetMealPlanByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mealPlan)
}

// UpdateMealPlan handles the update of a meal plan by ID
func (h *MealPlanHandler) UpdateMealPlan(c *gin.Context) {
	id, err := utils.ConvertParamToInt(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedMealPlan models.MealPlan

	if err := c.ShouldBindJSON(&updatedMealPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedMealPlan, err2 := h.service.UpdateMealPlan(id, updatedMealPlan)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err2.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedMealPlan)
}

// DeleteMealPlan handles the deletion of a meal plan by ID
func (h *MealPlanHandler) DeleteMealPlan(c *gin.Context) {
	id, err := utils.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err2 := h.service.DeleteMealPlan(id)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err2.Error()})
		return
	}

	c.Status(http.StatusOK)
}
