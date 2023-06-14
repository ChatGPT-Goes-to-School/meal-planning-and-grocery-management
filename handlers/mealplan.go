package handlers

import (
	"net/http"
	"strconv"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/repositories"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/services"
	"github.com/gin-gonic/gin"
)

// MealPlanHandler handles the HTTP requests for meal plans
type MealPlanHandler struct {
	service    services.MealPlanService
	repository repositories.MealPlanRepository
}

// CreateMealPlan handles the creation of a meal plan
func (h *MealPlanHandler) CreateMealPlan(c *gin.Context) {
	var mealPlan models.MealPlan
	if err := c.ShouldBindJSON(&mealPlan); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMealPlan := h.service.CreateMealPlan(mealPlan)
	h.repository.CreateMealPlan(createdMealPlan)

	c.JSON(http.StatusCreated, createdMealPlan)
}

// GetMealPlan handles the retrieval of a meal plan by ID
func (h *MealPlanHandler) GetMealPlan(c *gin.Context) {
	id, err := h.ConvertParamToInt(c.Param("id"))

	mealPlan, err := h.service.GetMealPlan(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mealPlan)
}

// UpdateMealPlan handles the update of a meal plan by ID
func (h *MealPlanHandler) UpdateMealPlan(c *gin.Context) {
	id, err := h.ConvertParamToInt(c.Param("id"))

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

	h.repository.UpdateMealPlan(id, updatedMealPlan)

	c.JSON(http.StatusOK, updatedMealPlan)
}

// DeleteMealPlan handles the deletion of a meal plan by ID
func (h *MealPlanHandler) DeleteMealPlan(c *gin.Context) {
	id, err := h.ConvertParamToInt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err2 := h.service.DeleteMealPlan(id)
	if err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err2.Error()})
		return
	}

	h.repository.DeleteMealPlan(id)

	c.Status(http.StatusNoContent)
}

func (h *MealPlanHandler) ConvertParamToInt(param string) (int, error) {
    id, err := strconv.Atoi(param)
    if err != nil {
        return 0, err
    }
    return id, nil
}