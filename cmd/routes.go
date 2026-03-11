package cmd

import (
	"net/http"
	"github.com/luminous479/product-list/handler"
)

func initRoutes(mux *http.ServeMux) {
	mux.Handle("GET /hello", http.HandlerFunc(handler.TestHandler))
	mux.Handle("GET /products", http.HandlerFunc(handler.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProductHandler))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handler.GetProductByID))
}
