package product

import (
	"net/http"
	"github.com/luminous479/product-list/utils"
)

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {

	if data, err := h.svc.List(); err == nil {
		utils.SendData(w, data, http.StatusOK)
		return
	}
	utils.SendData(w, "No products found", http.StatusNotFound)

}
