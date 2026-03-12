package rest

import (
	"net/http"

	"github.com/luminous479/product-list/rest/handlers"
	"github.com/luminous479/product-list/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProductHandler)))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(handlers.GetProductByID)))
	mux.Handle("PUT /product/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /product/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))
}
