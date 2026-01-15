package handlers

import (
	"net/http"
	"github.com/TheAditya-10/go/phase5/repository"
)

type UserHandler struct {
	Repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{Repo: repo}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// For demonstration, using a static email
	email := "example@mail.com"
	user, err := h.Repo.GetUserByEmail(r.Context(), email)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	utils.JSON(w, http.StatusOK, user)
}