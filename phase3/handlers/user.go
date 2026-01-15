package handlers

import (
	"net/http"
	"github.com/TheAditya-10/go/phase3/models"
	"github.com/TheAditya-10/go/phase3/utils"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{
		ID:       1,
		Username: "Aditya",
		Email:    "aditya@example.com",
	}

	utils.JSON(w, http.StatusOK, user)
}