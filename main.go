package main

import (
	"fmt"
	"net/http"
	"encoding/json"
) 

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
var products  []Product
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	encoder := json.NewEncoder(w)
	err := encoder.Encode(products)
	if err != nil {
		http.Error(w, "Error encoding products", http.StatusInternalServerError)
		return
	}

}

func main() {
	// Create a new router
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/getProducts",getProducts)
	// Start the server
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}

func init() {
	products = []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
		{ID: 3, Name: "Product 3", Price: 5.49},
	}
	
}