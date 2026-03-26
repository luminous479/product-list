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
	totalCount, err := h.svc.Count()
	if err != nil {
		utils.SendData(w, err, http.StatusConflict)
	}

	data, error := h.svc.List(page, limit)
	pagination := utils.Pagination{
		Data: data,
		PagenatedData: utils.PagenatedData{
			Page:       int64(page),
			Limit:      int64(limit),
			TotalItems: totalCount,
			TotalPages: int64(totalCount / 20),
		},
	}
	if error == nil {
		utils.SendPaginatedData(w, pagination)
		return
	}

	utils.SendData(w, "No products found", http.StatusNotFound)

}
