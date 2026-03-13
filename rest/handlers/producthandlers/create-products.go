package producthandlers

import (
	"encoding/json"
	"net/http"

	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	var newProduct database.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}

	product := database.AddProduct(newProduct)
	w.WriteHeader(http.StatusCreated)
	utils.SendData(w, product, http.StatusCreated)
}
