package models

import (
	"time"

	"gorm.io/gorm"
)

type Discard struct {
	gorm.Model
	UserID   uint      `json:"user_id"`
	City     string    `json:"city"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}
