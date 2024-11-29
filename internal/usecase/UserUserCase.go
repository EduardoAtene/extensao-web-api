package usecase

import (
	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"
	"github.com/EduardoAtene/extensao-web-api/internal/repository"
	"gorm.io/gorm"
)

type UserUserCase struct {
	userRepo *repository.UserRepository
}

func NewUserUserCase(db *gorm.DB) *UserUserCase {
	return &UserUserCase{
		userRepo: repository.NewUserRepository(db),
	}
}

func (uc *UserUserCase) GetUserDetails(userId uint) (*models.User, error) {
	return uc.userRepo.FindByID(userId)
}
