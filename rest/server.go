package rest

import (
	"fmt"
	"net/http"

	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/rest/handlers/product"
	"github.com/luminous479/product-list/rest/handlers/user"
	"github.com/luminous479/product-list/rest/middlewares"
)

type Server struct {
	productHandler *product.ProductHandler
	userHandler    *user.UserHandler
	conf config.Config
}

func NewServer( productHandler *product.ProductHandler, 
	userHandler *user.UserHandler, config config.Config) *Server {	
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		conf:             config,
	}
}

func (s *Server) StartServer() {
	manager := middlewares.NewManager()
	manager.Use(middlewares.Cors, middlewares.Preflight, middlewares.Logger)
	mux := http.NewServeMux()

	// Register routes for product and user handlers
	s.productHandler.RegisterRoutes(mux, manager)
	s.userHandler.RegisterRoutes(mux, manager)

	// Start the server
	addr := fmt.Sprintf(":%d", s.conf.HttpPort)
	fmt.Println("Server is running on", addr)
	err := http.ListenAndServe(addr, manager.WrapMux(mux))
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
