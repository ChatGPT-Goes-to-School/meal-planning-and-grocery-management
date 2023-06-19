package main

import (
	"fmt"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/config"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	db, err := config.ConnectDB()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Close
	defer sqlDB.Close()

	// Initialize the Gin router
	router := gin.Default()

	// Register the API endpoints
	routes.SetupMealPlanRoutes(router, db)

	// Run the server
	router.Run(":8080")
}
