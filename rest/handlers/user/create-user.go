package user

import (
	"encoding/json"
	"net/http"
	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/utils"
)

type RequestCreateUser struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"pass"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var  Request RequestCreateUser
	err := json.NewDecoder(r.Body).Decode(&Request)	
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	user, err :=  h.userRepo.Create(repo.User{
		Name:		Request.Name,
		Email:		Request.Email,
		Password:	Request.Password,
		IsShopOwner: Request.IsShopOwner,
	})
	if err == nil {
		utils.SendData(w,user,http.StatusCreated)
		return
	} 
	utils.SendData(w, "Failed to create user", http.StatusInternalServerError)
	

}