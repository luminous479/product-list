package product

import (
	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/rest/middlewares"
)

type ProductHandler struct {
	middlewares *middlewares.Middlewares
	productRepo repo.ProductRepo
}

func NewProductHandler(middlewares *middlewares.Middlewares, productRepo repo.ProductRepo) *ProductHandler {
	return &ProductHandler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}