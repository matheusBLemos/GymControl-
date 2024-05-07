package web

import (
	"GusLem/gymControll/internal/dto"
	"GusLem/gymControll/internal/entity"
	"GusLem/gymControll/internal/usecase"
	"encoding/json"
	"net/http"
)

type WebUserHandler struct {
	UserInterface entity.UserInterface
}

func NewUserHandler(
	UserRepository entity.UserInterface,
) *WebUserHandler {
	return &WebUserHandler{
		UserInterface: UserRepository,
	}
}

func (h *WebUserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto dto.UserDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createUser := usecase.NewUser(h.UserInterface)
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