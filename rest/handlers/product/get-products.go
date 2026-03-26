package product

import (
	"net/http"
	"strconv"

	"github.com/luminous479/product-list/utils"
)

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {


	reqQuery := r.URL.Query()

pageStr := reqQuery.Get("page")
limitStr := reqQuery.Get("limit")

// defaults
page := int(1)
limit := int(20)

// parse page
if p, err := strconv.ParseInt(pageStr, 10, 64); err == nil && p > 0 {
	page = int(p)
}

// parse limit
if l, err := strconv.ParseInt(limitStr, 10, 64); err == nil && l > 0 {
	limit = int(l)
}

	if data, err := h.svc.List(page, limit); err == nil {
		utils.SendData(w, data, http.StatusOK)
		return
	}
	utils.SendData(w, "No products found", http.StatusNotFound)

}
