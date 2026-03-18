package product

import (

	"github.com/luminous479/product-list/rest/middlewares"
)

type ProductHandler struct {
	middlewares *middlewares.Middlewares
	svc Service
}

func NewProductHandler(svc Service, middlewares *middlewares.Middlewares) *ProductHandler {
	return &ProductHandler{
		svc: svc,
		middlewares: middlewares,
		
	}
}