package handlers

import (
	"net/http"

	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	newProduct := utils.DecodeProduct(r, w)
	if (newProduct == database.Product{}) {
		return
	}
	database.AddProduct(newProduct)
	w.WriteHeader(http.StatusCreated)
	utils.SendData(w, newProduct, http.StatusCreated)
}
