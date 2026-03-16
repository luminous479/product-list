package user

import (
	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/repo"
)

type UserHandler struct {
	userRepo repo.UserRepo
    config   *config.Config
}

func NewUserHandler(userRepo repo.UserRepo, config *config.Config) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
		config:   config,
	}
}