package cmd

import (
	"fmt"
	"os"

	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/infra/db"
	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/rest"
	"github.com/luminous479/product-list/rest/handlers/product"
	"github.com/luminous479/product-list/rest/handlers/user"
	"github.com/luminous479/product-list/rest/middlewares"
	_ "github.com/lib/pq"
)

func Serve() {
	config := config.GetConfig()
    
	dbCon, err := db.NewConnection(config.DB)

	if err != nil {

		fmt.Println(err)	
		os.Exit(1)	
	}

	middlewares := middlewares.NewMiddlewares(config)
	productRepo := repo.NewProductRepo(dbCon)
	productHandler := product.NewProductHandler(middlewares, productRepo)
	userRepo := repo.NewUserRepo(dbCon)
	userHandler := user.NewUserHandler(userRepo, config)
	rest.NewServer(productHandler, userHandler, *config).StartServer()

}
