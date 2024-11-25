package usecase

import (
	"time"

	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"
	"github.com/EduardoAtene/extensao-web-api/internal/repository"
	"gorm.io/gorm"
)

type DiscardUseCase struct {
	discardRepo  *repository.DiscardRepository
	rewardRepo   *repository.RewardRepository
	locationRepo *repository.LocationRepository
}

type CreateDiscardRequest struct {
	City     string    `json:"city"`
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
}

func NewDiscardUseCase(db *gorm.DB) *DiscardUseCase {
	return &DiscardUseCase{
		discardRepo:  repository.NewDiscardRepository(db),
		rewardRepo:   repository.NewRewardRepository(db),
		locationRepo: repository.NewLocationRepository(db),
	}
}

func (uc *DiscardUseCase) CreateDiscard(userID uint, req CreateDiscardRequest) ([]models.UserReward, error) {
	// Criar o descarte
	discard := models.Discard{
		UserID:   userID,
		City:     req.City,
		Location: req.Location,
		Date:     req.Date,
	}

	if err := uc.discardRepo.Create(&discard); err != nil {
		return nil, err
	}

	// Atualizar local único
	if err := uc.locationRepo.IncrementLocation(req.City, req.Location); err != nil {
		return nil, err
	}

	// Verificar e atribuir recompensas
	return uc.checkAndAwardRewards(userID)
}

func (uc *DiscardUseCase) checkAndAwardRewards(userID uint) ([]models.UserReward, error) {
	var newRewards []models.UserReward

	// Contar total de descartes
	totalDiscards, err := uc.discardRepo.CountUserDiscards(userID)
	if err != nil {
		return nil, err
	}

	// Contar locais únicos
	uniqueLocations, err := uc.locationRepo.CountUniqueLocations(userID)
	if err != nil {
		return nil, err
	}

	// Verificar cada tipo de recompensa
	rewards, err := uc.rewardRepo.GetAllRewards()
	if err != nil {
		return nil, err
	}

	for _, reward := range rewards {
		hasReward, err := uc.rewardRepo.UserHasReward(userID, reward.ID)
		if err != nil {
			return nil, err
		}

		if hasReward {
			continue
		}

		shouldAward := false

		switch reward.Name {
		case "Bronze":
			shouldAward = totalDiscards >= 10
		case "Viajante":
			shouldAward = uniqueLocations >= 5
		case "Prata":
			shouldAward = totalDiscards >= 20
		case "Ouro":
			shouldAward = totalDiscards >= 30
		case "Variado":
			shouldAward = uniqueLocations >= 20
		case "Diamante":
			shouldAward = totalDiscards >= 50
		}

		if shouldAward {
			userReward := models.UserReward{
				UserID:     userID,
				RewardID:   reward.ID,
				AchievedAt: time.Now(),
			}

			if err := uc.rewardRepo.CreateUserReward(&userReward); err != nil {
				return nil, err
			}

			newRewards = append(newRewards, userReward)
		}
	}

	return newRewards, nil
}

func (uc *DiscardUseCase) GetUserRewards(userID uint) ([]models.UserReward, error) {
	return uc.rewardRepo.GetUserRewards(userID)
}
