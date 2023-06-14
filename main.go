package main

import (
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize the Gin router
	router := gin.Default()


	// Register the API endpoints
	routes.SetupRoutes(router)

	// Run the server
	router.Run(":8080")
}
