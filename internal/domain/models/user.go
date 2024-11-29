package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name" gorm:"default:'Usuário'"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	Phone     string    `json:"phone"`
}
