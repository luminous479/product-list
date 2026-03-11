package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}
	database.AddProduct(newProduct)	
	w.WriteHeader(http.StatusCreated)
	utils.SendData(w, newProduct, http.StatusCreated)
}
