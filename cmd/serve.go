package cmd

import (
	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/rest"
	"github.com/luminous479/product-list/rest/handlers/product"
	"github.com/luminous479/product-list/rest/handlers/user"
	"github.com/luminous479/product-list/rest/middlewares"
)

func Serve() {
	config := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(config)
	productHandler := product.NewProductHandler(middlewares)
	userHandler := user.NewUserHandler()
	rest.NewServer(productHandler, userHandler, *config).StartServer()

}
