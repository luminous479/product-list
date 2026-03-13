package rest

import (
	"fmt"
	"net/http"
	"github.com/luminous479/product-list/env"
	"github.com/luminous479/product-list/rest/middlewares/product"

)
func StartServer( config env.Config) {
	manager := product.NewManager()
	manager.Use(product.Cors, product.Preflight, product.Logger)
	mux := http.NewServeMux()
	initRoutes(mux, manager)
	// Start the server
    addr := fmt.Sprintf(":%d", config.HttpPort)
	fmt.Println("Server is running on", addr)
	err := http.ListenAndServe(addr, manager.WrapMux(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}