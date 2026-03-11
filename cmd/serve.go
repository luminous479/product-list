package cmd

import (
	"fmt"
	"net/http"
	"github.com/luminous479/product-list/middleware"
)

func Serve() {
	manager := middleware.NewManager()
	manager.Use(middleware.Cors, middleware.Preflight, middleware.Logger)
	mux := http.NewServeMux()
	initRoutes(mux)
	// Start the server
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", manager.WrapMux(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
