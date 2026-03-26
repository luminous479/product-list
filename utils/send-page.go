package utils

import (
	"net/http"
	"github.com/luminous479/product-list/domain"
)

type Pagination struct {
	Data          []*domain.Product `json:"data"`
	PagenatedData PagenatedData     `json:pafinatedData`
}

type PagenatedData struct {
	Page       int64 `json:page`
	Limit      int64 `json:limit`
	TotalItems int64 `json:totalItems`
	TotalPages int64 `json:totalPages`
}

func SendPaginatedData(w http.ResponseWriter, date Pagination) {
	SendData(w, date, http.StatusOK)
}
