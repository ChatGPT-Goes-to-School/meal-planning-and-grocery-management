package models

import "gorm.io/gorm"

type Grocery struct {
	gorm.Model
	Name     string     `json:"name"`
	Username string     `json:"username"`
	Items    []ListItem `gorm:"foreignKey:GroceryID"`
}
