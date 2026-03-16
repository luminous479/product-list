package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/utils"
)

type RequestUpdateProduct struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	// Extract the product ID from the URL path
	productId := r.PathValue("id")
	// Convert the ID string to an integer
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	// Decode the request body into a Product struct
	var req RequestUpdateProduct
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}

	updatedProduct := repo.Product{
		ID:    id,
		Name:  req.Name,
		Price: req.Price,
	}

	// Update the product in the database

	 data := h.productRepo.Update(updatedProduct)
	 if data == nil {
		utils.SendData(w, "Invalid product credentials", http.StatusNotFound)
		return
	 }
		utils.SendData(w, data, http.StatusOK)
	

}
