package handlers

import (
	
	"net/http"
	"strconv"

	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path
	productId := r.PathValue("id")
	// Convert the ID string to an integer
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// Decode the request body into a Product struct
	updatedProduct := utils.DecodeProduct(r, w)
	if (updatedProduct == database.Product{}) {
		return 
	}
	updatedProduct.ID = id
	// Update the product in the database
	if data := database.UpdateProduct(id, updatedProduct); data != (database.Product{}) {
		utils.SendData(w, data, http.StatusOK)
		return
	}
	utils.SendData(w, "Product not found", http.StatusNotFound)
}
