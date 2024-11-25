package repository

import (
	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"

	"gorm.io/gorm"
)

type DiscardRepository struct {
	db *gorm.DB
}

func NewDiscardRepository(db *gorm.DB) *DiscardRepository {
	return &DiscardRepository{db: db}
}

func (r *DiscardRepository) Create(discard *models.Discard) error {
	return r.db.Create(discard).Error
}

func (r *DiscardRepository) CountUserDiscards(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Discard{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}
