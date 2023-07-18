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

// MealPlanHandler handles the HTTP requests for meal plans
type MealPlanHandler struct {
	service *services.MealPlanService
	db      *gorm.DB
}

func NewMealPlanHandler(db *gorm.DB) *MealPlanHandler {
	service := services.NewMealPlanService(db)
	return &MealPlanHandler{
		service: service,
		db:      db,
	}
}

// CreateMealPlan creates a new meal plan.
// @Summary Create a meal plan
// @Description Create a new meal plan
// @Tags meal-plans
// @Accept json
// @Produce json
// @Param mealPlan body models.MealPlan true "Meal plan object"
// @Success 201 {object} models.MealPlan
// @Failure 400 {object} httputil.HTTPError
// @Router /meal-plans [post]
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

// GetMealPlan retrieves a meal plan by ID.
// @Summary Get a meal plan by ID
// @Description Retrieve a meal plan by its ID
// @Tags meal-plans
// @Accept json
// @Produce json
// @Param id path int true "Meal plan ID"
// @Success 200 {object} models.MealPlan
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /meal-plans/{id} [get]
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

// GetMealPlanByUsername retrieves a meal plan by username.
// @Summary Get a meal plan by username
// @Description Retrieve a meal plan by the username of the owner
// @Tags meal-plans
// @Accept json
// @Produce json
// @Param username path string true "Username of the meal plan owner"
// @Success 200 {object} models.MealPlan
// @Failure 404 {object} httputil.HTTPError
// @Router /meal-plans/username/{username} [get]
func (h *MealPlanHandler) GetMealPlanByUsername(c *gin.Context) {
	username := c.Param("username")

	mealPlan, err := h.service.GetMealPlanByUsername(username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mealPlan)
}

// UpdateMealPlan updates a meal plan.
// @Summary Update a meal plan
// @Description Update an existing meal plan by its ID
// @Tags meal-plans
// @Accept json
// @Produce json
// @Param id path int true "ID of the meal plan to update"
// @Param updatedMealPlan body models.MealPlan true "Updated meal plan object"
// @Success 200 {object} models.MealPlan
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /meal-plans/{id} [put]
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

// DeleteMealPlan deletes a meal plan.
// @Summary Delete a meal plan
// @Description Delete an existing meal plan by its ID
// @Tags meal-plans
// @Param id path int true "ID of the meal plan to delete"
// @Success 200 "OK"
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Router /meal-plans/{id} [delete]
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

func (h *MealPlanHandler) GetMealPlanByUsernameAndDatePlan(c *gin.Context) {
	username := c.Param("username")
	datePlan := c.Param("datePlan")

	mealPlan, err := h.service.GetMealPlanByUsernameAndDatePlan(username, datePlan)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mealPlan)
}
