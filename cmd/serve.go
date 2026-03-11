package cmd

import (
	"github.com/luminous479/product-list/env"
	"github.com/luminous479/product-list/rest"
)

func Serve() {
	config := env.GetConfig()
	rest.StartServer(config)

}
