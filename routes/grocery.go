package routes

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupGroceryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	groceryHandler := handlers.NewGroceryHandler(db)
	// Initialize the handlers
	router.GET(":id", groceryHandler.GetGrocery)
	router.GET("username/:username", groceryHandler.GetGroceriesByUsername)
	router.POST("", groceryHandler.CreateGrocery)
	router.PUT(":id", groceryHandler.UpdateGrocery)
	router.DELETE(":id", groceryHandler.DeleteGrocery)
	// Define more routes as needed
}
