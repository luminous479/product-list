package middlewares

import "github.com/luminous479/product-list/config"

type Middlewares struct {
	conf *config.Config
}

func  NewMiddlewares(conf *config.Config) *Middlewares {
	return &Middlewares{
		conf: conf,
	}
}
