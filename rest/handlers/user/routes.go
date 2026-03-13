package user

import (
	"net/http"

	"github.com/luminous479/product-list/rest/middlewares"
)

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /users/register", manager.With(http.HandlerFunc(h.CreateUser)))
	mux.Handle("GET /users/login", manager.With(http.HandlerFunc(h.LoginUser)))
}
