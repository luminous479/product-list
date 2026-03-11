package cmd

import (
	"fmt"
	"net/http"
	"github.com/luminous479/product-list/middleware"
	"github.com/luminous479/product-list/router"
)

func Serve() {
	manager := middleware.NewManager()
	mux := http.NewServeMux()
	manager.Use(middleware.Logger)
	allRoutes(mux, manager)
	router := router.GlobalRouter(mux)
	// Start the server
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
