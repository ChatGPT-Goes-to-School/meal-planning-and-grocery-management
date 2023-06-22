package main

import (
	"fmt"

	"os"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/config"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/docs"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/middlewares"
	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	docs.SwaggerInfo.Title = "Meal Planning and Grocery Management API"
	docs.SwaggerInfo.Description = "An API for managing meal plans and grocery lists."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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

	// Apply the JWT middleware to routes that require authentication
	router.Use(middlewares.JWTMiddleware(os.Getenv("JWT_SECRET")))
	// Register the API endpoints
	apiRoutes := router.Group(docs.SwaggerInfo.BasePath)
	{
		mealPlan := apiRoutes.Group("/meal-plans")
		{
			routes.SetupMealPlanRoutes(mealPlan, db)
		}
		grocery := apiRoutes.Group("/groceries")
		{
			routes.SetupGroceryRoutes(grocery, db)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Run the server
	router.Run(":8080")
}
