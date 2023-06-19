package routes

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMealPlanRoutes(router *gin.Engine, db *gorm.DB) {
	mealPlanHandler := handlers.NewMealPlanHandler(db)
	// Initialize the handlers
	router.POST("/meal-plans", mealPlanHandler.CreateMealPlan)
	router.GET("/meal-plans/username/:username", mealPlanHandler.GetMealPlanByUsername)
	router.GET("/meal-plans/:id", mealPlanHandler.GetMealPlan)
	router.PUT("/meal-plans/:id", mealPlanHandler.UpdateMealPlan)
	router.DELETE("/meal-plans/:id", mealPlanHandler.DeleteMealPlan)
	// Define more routes as needed
}
