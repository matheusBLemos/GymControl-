package controller

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
	"GusLem/gymControll/internal/usecase"
	"encoding/json"
	"net/http"
)

type WebOrderHandler struct {
	UserRepositoryWebHandler entity.UserInterface
}

func NewWebOrderHandler(
	UserRepository entity.UserInterface,
) *WebOrderHandler {
	return &WebOrderHandler{
		UserRepositoryWebHandler: UserRepository,
	}
}

func (h *WebOrderHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createUser := usecase.NewUser(h.UserRepositoryWebHandler)
	output, err := createUser.CreateNewUser(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}