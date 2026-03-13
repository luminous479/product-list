package product

import "github.com/luminous479/product-list/rest/middlewares"

type ProductHandler struct {
	middlewares *middlewares.Middlewares
}

func NewProductHandler(middlewares *middlewares.Middlewares) *ProductHandler {
	return &ProductHandler{
		middlewares: middlewares, 
	}
}