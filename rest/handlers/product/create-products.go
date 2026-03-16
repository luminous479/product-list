package product

import (
	"encoding/json"
	"net/http"

	"github.com/luminous479/product-list/repo"
	"github.com/luminous479/product-list/utils"
)

type RequestCreateProduct struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var req RequestCreateProduct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}

	product, err := h.productRepo.Create(repo.Product{
		Name:  req.Name,
		Price: req.Price,
	})
	if err != nil {
		http.Error(w, "Error creating product", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	utils.SendData(w, product, http.StatusCreated)
}
