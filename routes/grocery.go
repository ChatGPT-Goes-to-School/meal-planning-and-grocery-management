package routes

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGroceryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	groceryHandler := handlers.NewGroceryHandler(db)
	// Initialize the handlers
	router.GET("/groceries/:id", groceryHandler.GetGrocery)
	router.GET("/groceries/username/:username", groceryHandler.GetGroceriesByUsername)
	router.POST("/groceries", groceryHandler.CreateGrocery)
	router.PUT("/groceries/:id", groceryHandler.UpdateGrocery)
	router.DELETE("/groceries/:id", groceryHandler.DeleteGrocery)
	// Define more routes as needed
}
