package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoAtene/extensao-web-api/internal/usecase"
	"github.com/EduardoAtene/extensao-web-api/internal/utils"
	"gorm.io/gorm"
)

type DiscardHandler struct {
	discardUseCase *usecase.DiscardUseCase
}

func NewDiscardHandler(db *gorm.DB) *DiscardHandler {
	return &DiscardHandler{
		discardUseCase: usecase.NewDiscardUseCase(db),
	}
}

func (h *DiscardHandler) CreateDiscard(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	var req usecase.CreateDiscardRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro ao processar a requisição", err.Error()), http.StatusBadRequest)
		return
	}

	rewards, err := h.discardUseCase.CreateDiscard(userID, req)
	if err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro ao registrar o descarte", err.Error()), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"message":     "Descarte registrado com sucesso",
		"new_rewards": rewards,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *DiscardHandler) GetUserRewards(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(uint)

	rewards, err := h.discardUseCase.GetUserRewards(userID)
	if err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro ao buscar recompensas", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rewards)
}
