package user

import (
	"github.com/luminous479/product-list/config"
)

type UserHandler struct {
    config   *config.Config
	svc Service
}

func NewUserHandler(svc Service, config *config.Config) *UserHandler {
	return &UserHandler{
		svc: svc,
		config:   config,
	}
}