package cmd

import (
	"fmt"
	"net/http"
	"github.com/luminous479/product-list/handler"
	"github.com/luminous479/product-list/router"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /products", http.HandlerFunc(handler.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handler.CreateProductHandler))
	mux.Handle("GET /products/{productId}", http.HandlerFunc(handler.GetProductByID))
	router := router.GlobalRouter(mux)

	// Start the server
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
