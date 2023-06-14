package routes

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	mealPlanHandler := handlers.MealPlanHandler{}
	// Initialize the handlers
	router.POST("/meal-plans", mealPlanHandler.CreateMealPlan)
	router.GET("/meal-plans/:id", mealPlanHandler.GetMealPlan)
	router.PUT("/meal-plans/:id", mealPlanHandler.UpdateMealPlan)
	router.DELETE("/meal-plans/:id", mealPlanHandler.DeleteMealPlan)
	// Define more routes as needed
}