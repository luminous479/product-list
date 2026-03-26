package cmd

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/luminous479/product-list/config"
	"github.com/luminous479/product-list/infra/db"
	"github.com/luminous479/product-list/product"
	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/rest"
	productHandler "github.com/luminous479/product-list/rest/handlers/product"
	UserHandler "github.com/luminous479/product-list/rest/handlers/user"
	"github.com/luminous479/product-list/rest/middlewares"
	"github.com/luminous479/product-list/user"
)

func Serve() {
	config := config.GetConfig()

	dbCon, err := db.NewConnection(config.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db.DBMigrate(dbCon, "./migrations")
	middlewares := middlewares.NewMiddlewares(config)
	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)
	//domains
	userSvc := user.NewService(userRepo)
	productsvc := product.NewService(productRepo)
	productHandler := productHandler.NewProductHandler(productsvc, middlewares)
	userHandler := UserHandler.NewUserHandler(userSvc, config)
	rest.NewServer(productHandler, userHandler, *config).StartServer()

}
