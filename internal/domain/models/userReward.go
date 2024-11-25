package models

import (
	"time"

	"gorm.io/gorm"
)

type UserReward struct {
	gorm.Model
	UserID     uint       `json:"user_id"`
	RewardID   uint       `json:"reward_id"`
	Reward     RewardType `json:"reward"`
	AchievedAt time.Time  `json:"achieved_at"`
}
