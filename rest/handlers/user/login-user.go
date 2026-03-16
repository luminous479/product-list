package user

import (
	"encoding/json"
	"net/http"
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
	user, err := h.userRepo.Find(loginReq.Email)
	if err != nil {
		utils.SendData(w, "User not found", http.StatusUnauthorized)
		return
	}

	accessToken, err := utils.CreateJwt(h.config.JwtSecretKey, utils.PayLoad{
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
