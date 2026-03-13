package userhandlers

import (
	"encoding/json"
	"net/http"
	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user database.User
	err := json.NewDecoder(r.Body).Decode(&user)	
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	createdUser := database.CreateUser(user)
	utils.SendData(w,createdUser,http.StatusCreated)
}