package user

import (
	"encoding/json"
	"net/http"

	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/utils"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"pass"`
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginReq loginRequest
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	user := database.GetUserByEmail(loginReq.Email)
	if user == nil || user.Password != loginReq.Password {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	conf := config.GetConfig()

	accessToken, err := utils.CreateJwt(conf.JwtSecretKey, utils.PayLoad{
		Sub:         user.ID,
		Name:        user.Name,
		Email:       user.Email,
		IsShopOwner: user.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Error creating access token", http.StatusInternalServerError)
		return
	}
	utils.SendData(w, accessToken, http.StatusOK)
}
