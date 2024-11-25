package repository

import (
	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"

	"gorm.io/gorm"
)

type LocationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) *LocationRepository {
	return &LocationRepository{db: db}
}

func (r *LocationRepository) IncrementLocation(city, location string) error {
	var loc models.Location
	result := r.db.Where("city = ? AND location = ?", city, location).FirstOrCreate(&loc, models.Location{
		City:     city,
		Location: location,
		Count:    0,
	})
	if result.Error != nil {
		return result.Error
	}

	return r.db.Model(&loc).Update("count", loc.Count+1).Error
}

func (r *LocationRepository) CountUniqueLocations(userID uint) (int64, error) {
	var count int64
	err := r.db.Model(&models.Discard{}).
		Select("COUNT(DISTINCT CONCAT(city, location))").
		Where("user_id = ?", userID).
		Count(&count).Error
	return count, err
}
