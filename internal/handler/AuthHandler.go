package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/EduardoAtene/extensao-web-api/internal/usecase"
	"github.com/EduardoAtene/extensao-web-api/internal/utils"
	"gorm.io/gorm"
)

type AuthHandler struct {
	authUseCase *usecase.AuthUseCase
}

func NewAuthHandler(db *gorm.DB) *AuthHandler {
	return &AuthHandler{
		authUseCase: usecase.NewAuthUseCase(db),
	}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req usecase.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro ao processar a requisição", err.Error()), http.StatusBadRequest)
		return
	}

	if err := h.authUseCase.Register(req); err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro ao criar usuário", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário criado com sucesso"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req usecase.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro ao processar a requisição", err.Error()), http.StatusBadRequest)
		return
	}

	token, err := h.authUseCase.Login(req)
	if err != nil {
		http.Error(w, utils.GenerateErrorResponse("Erro de autenticação", err.Error()), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
