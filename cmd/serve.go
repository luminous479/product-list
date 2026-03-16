package cmd

import (
	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/rest"
	"github.com/luminous479/product-list/rest/handlers/product"
	"github.com/luminous479/product-list/rest/handlers/user"
	"github.com/luminous479/product-list/rest/middlewares"
)

func Serve() {
	config := config.GetConfig()
	middlewares := middlewares.NewMiddlewares(config)
	productRepo := repo.NewProductRepo()
	productHandler := product.NewProductHandler(middlewares, productRepo)
	userRepo := repo.NewUserRepo()
	userHandler := user.NewUserHandler(userRepo, config)
	rest.NewServer(productHandler, userHandler, *config).StartServer()

}
