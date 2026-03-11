package cmd

import (
	"net/http"
	"github.com/luminous479/product-list/handler"
	"github.com/luminous479/product-list/middleware"
)

func allRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /hello", manager.With(http.HandlerFunc(handler.TestHandler)))
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handler.CreateProductHandler)))
	mux.Handle("GET /products/{productId}", manager.With(http.HandlerFunc(handler.GetProductByID)))
}
