package usecase

import (
	"time"

	"github.com/EduardoAtene/extensao-web-api/internal/domain/models"
	"github.com/EduardoAtene/extensao-web-api/internal/repository"
	"github.com/EduardoAtene/extensao-web-api/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthUseCase struct {
	userRepo *repository.UserRepository
}

func NewAuthUseCase(db *gorm.DB) *AuthUseCase {
	return &AuthUseCase{
		userRepo: repository.NewUserRepository(db),
	}
}

type RegisterRequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (uc *AuthUseCase) Register(req RegisterRequest) error {
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.User{
		Name:      req.Name,
		Email:     req.Email,
		Password:  string(hashedPassword),
		BirthDate: birthDate,
		Phone:     req.Phone,
	}

	return uc.userRepo.Create(&user)
}

func (uc *AuthUseCase) Login(req LoginRequest) (string, error) {
	user, err := uc.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}

	return utils.GenerateJWT(user.ID)
}
