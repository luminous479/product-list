package rest

import (
	"net/http"
	"github.com/luminous479/product-list/rest/handlers"
)

func initRoutes(mux *http.ServeMux) {
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProductHandler))
	mux.Handle("GET /product/{id}", http.HandlerFunc(handlers.GetProductByID))
}
