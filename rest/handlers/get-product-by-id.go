package handlers

import (
	"net/http"
	"strconv"
	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path
	productId := r.PathValue("productId")
	// Convert the ID string to an integer
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// Search for the product in the database
	for _, product := range database.Products {
		if product.ID == id {
			utils.SendData(w, product, http.StatusOK)
			return
		}
	}
	utils.SendData(w, "Product not found", http.StatusNotFound)

}
