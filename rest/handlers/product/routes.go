package product

import (
	"net/http"

	"github.com/luminous479/product-list/rest/middlewares"
)

func (h *ProductHandler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProductHandler),
		h.middlewares.JwtAuthentication))
	mux.Handle("GET /product/{id}", manager.With(http.HandlerFunc(h.GetProductByID)))
	mux.Handle("PUT /product/{id}", manager.With(http.HandlerFunc(h.UpdateProduct), h.middlewares.JwtAuthentication))
	mux.Handle("DELETE /product/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), h.middlewares.JwtAuthentication))

}
