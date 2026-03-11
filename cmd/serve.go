package cmd

import (
	"fmt"
	"net/http"

	"github.com/luminous479/product-list/env"
	"github.com/luminous479/product-list/middleware"
)

func Serve() {
	config := env.GetConfig()
	manager := middleware.NewManager()
	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)
	mux := http.NewServeMux()
	initRoutes(mux)
	// Start the server
    addr := fmt.Sprintf(":%d", config.HttpPort)
	fmt.Println("Server is running on", addr)
	err := http.ListenAndServe(addr, manager.WrapMux(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
