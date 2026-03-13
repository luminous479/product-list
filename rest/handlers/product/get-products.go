package product

import (
	"net/http"
	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.GetProducts(), http.StatusOK)
}
