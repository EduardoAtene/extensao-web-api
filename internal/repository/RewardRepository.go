package repository

import (
	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"

	"gorm.io/gorm"
)

type RewardRepository struct {
	db *gorm.DB
}

func NewRewardRepository(db *gorm.DB) *RewardRepository {
	return &RewardRepository{db: db}
}

func (r *RewardRepository) GetAllRewards() ([]models.RewardType, error) {
	var rewards []models.RewardType
	err := r.db.Find(&rewards).Error
	return rewards, err
}

func (r *RewardRepository) UserHasReward(userID, rewardID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserReward{}).
		Where("user_id = ? AND reward_id = ?", userID, rewardID).
		Count(&count).Error
	return count > 0, err
}

func (r *RewardRepository) CreateUserReward(reward *models.UserReward) error {
	return r.db.Create(reward).Error
}

func (r *RewardRepository) GetUserRewards(userID uint) ([]models.UserReward, error) {
	var rewards []models.UserReward
	err := r.db.Preload("Reward").
		Where("user_id = ?", userID).
		Find(&rewards).Error
	return rewards, err
}
