package handler

import (
	"net/http"
	"github.com/luminous479/product-list/database"
	"github.com/luminous479/product-list/utils"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	utils.SendData(w, database.Products, http.StatusOK)
}
