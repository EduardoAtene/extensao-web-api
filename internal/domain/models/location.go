package models

import (
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	City     string `json:"city" gorm:"uniqueIndex:idx_city_location"`
	Location string `json:"location" gorm:"uniqueIndex:idx_city_location"`
	Count    int    `json:"count"`
}
