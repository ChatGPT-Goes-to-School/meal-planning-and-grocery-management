package routes

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupMealPlanRoutes(router *gin.RouterGroup, db *gorm.DB) {
	mealPlanHandler := handlers.NewMealPlanHandler(db)

	// Initialize the handlers
	router.POST("", mealPlanHandler.CreateMealPlan)
	router.GET("username/:username", mealPlanHandler.GetMealPlanByUsername)
	router.GET(":id", mealPlanHandler.GetMealPlan)
	router.PUT(":id", mealPlanHandler.UpdateMealPlan)
	router.DELETE(":id", mealPlanHandler.DeleteMealPlan)
	// Define more routes as needed
}
