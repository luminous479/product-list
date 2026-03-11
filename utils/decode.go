package utils

import (
	"encoding/json"
	"net/http"

	"github.com/luminous479/product-list/database"
)

func DecodeProduct( r *http.Request, w http.ResponseWriter) *database.Product {

	var updatedProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return &database.Product{} 
	}
	return &updatedProduct

}
