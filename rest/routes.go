package rest

import (
	"net/http"

	"github.com/luminous479/product-list/rest/handlers/producthandlers"
	"github.com/luminous479/product-list/rest/handlers/userhandlers"
	jwtauthentication "github.com/luminous479/product-list/rest/middlewares/jwt-authentication"
	"github.com/luminous479/product-list/rest/middlewares/product"
)

func initRoutes(mux *http.ServeMux, manager *product.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(producthandlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(producthandlers.CreateProductHandler), 
	jwtauthentication.JwtAuthenticationMiddleware))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(producthandlers.GetProductByID)))
	mux.Handle("PUT /product/{id}", manager.With(http.HandlerFunc(producthandlers.UpdateProduct),jwtauthentication.JwtAuthenticationMiddleware))
	mux.Handle("DELETE /product/{id}", manager.With(http.HandlerFunc(producthandlers.DeleteProduct),jwtauthentication.JwtAuthenticationMiddleware))
	// User routes
	mux.Handle("POST /users/register", manager.With(http.HandlerFunc(userhandlers.CreateUser)))
	mux.Handle("GET /users/login", manager.With(http.HandlerFunc(userhandlers.LoginUser)))

}
