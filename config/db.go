package config

import (
	"fmt"
	"os"

	"github.com/ChatGPT-Goes-to-School/meal-planning-and-grocery-management/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	// Get the environment variables
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")


	
	// Create the database connection string
	dbConnStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", username, password, host, port, dbName)
	
	// Set up the MySQL database connection string
	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})

	// Migrate the database
	db.AutoMigrate(&models.MealPlan{}, &models.Meal{})

	if err != nil {
		return nil, err
	}

	// Check the database connection
	dbase, err := db.DB()
	if err != nil {
		return nil, err
	}

	err = dbase.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the MySQL database!")

	return db, nil
}
