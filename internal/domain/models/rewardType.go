package models

import "gorm.io/gorm"

type RewardType struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Requirement int    `json:"requirement"` // Número de descartes necessários
}
