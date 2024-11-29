package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/EduardoAtene/extensao-web-api/internal/usecase"
	"github.com/EduardoAtene/extensao-web-api/internal/utils"
	"gorm.io/gorm"
)

type UserHandler struct {
	userUseCase *usecase.UserUserCase
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{
		userUseCase: usecase.NewUserUserCase(db),
	}
}

func (h *UserHandler) GetUserDetails(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	user, err := h.userUseCase.GetUserDetails(userID)
	if err != nil {
		http.Error(w, utils.GenerateErrorResponse("Usuário não encontrado", err.Error()), http.StatusNotFound)
		return
	}

	userResponse := struct {
		ID        uint      `json:"id"`
		Name      string    `json:"name"`
		Email     string    `json:"email"`
		BirthDate time.Time `json:"birth_date"`
		Phone     string    `json:"phone"`
	}{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		BirthDate: user.BirthDate,
		Phone:     user.Phone,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}
