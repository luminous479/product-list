package product

import (
	"net/http"
	"strconv"
	"github.com/luminous479/product-list/utils"
)

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	id, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}
	
	if data, err := h.productRepo.Delete(id); err == nil {
		utils.SendData(w, data, http.StatusOK)
		return
	}
	utils.SendData(w, "Product not found", http.StatusNotFound)
}
