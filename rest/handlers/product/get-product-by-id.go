package product

import (
	"net/http"
	"strconv"
	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path
	productId := r.PathValue("id")
	// Convert the ID string to an integer
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	if data := database.GetProductByID(id); data != (database.Product{}) {
		utils.SendData(w, database.GetProductByID(id), http.StatusOK)
		return
	}
	utils.SendData(w, "Product not found", http.StatusNotFound)

}
