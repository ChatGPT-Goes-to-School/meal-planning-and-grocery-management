package errors

import (
	"fmt"
)

type MealPlanNotFoundError struct {
    ID int
}

func (e MealPlanNotFoundError) Error() string {
    return fmt.Sprintf("meal plan not found with ID: %d", e.ID)
}

// Function to create a new instance of MealPlanNotFoundError
func NewMealPlanNotFoundError(id int) MealPlanNotFoundError {
    return MealPlanNotFoundError{ID: id}
}