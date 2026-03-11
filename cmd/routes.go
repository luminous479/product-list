package cmd

import (
	"net/http"
	"github.com/luminous479/product-list/handler"
	"github.com/luminous479/product-list/middleware"
)

func allRoutes(mux *http.ServeMux) {
	mux.Handle("GET /hello", middleware.NewManager().With(http.HandlerFunc(handler.TestHandler)))
	mux.Handle("GET /products", middleware.NewManager().With(http.HandlerFunc(handler.GetProducts)))
	mux.Handle("POST /products", middleware.NewManager().With(http.HandlerFunc(handler.CreateProductHandler)))
	mux.Handle("GET /products/{productId}", middleware.NewManager().With(http.HandlerFunc(handler.GetProductByID)))
}
