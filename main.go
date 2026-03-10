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
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newProduct Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)				
	if err != nil {
		http.Error(w, "Error decoding product", http.StatusBadRequest)
		return
	}
	newProduct.ID = len(products) + 1
	products = append(products, newProduct)
	w.WriteHeader(http.StatusCreated)	
	encoder := json.NewEncoder(w)
	err = encoder.Encode(newProduct)	
	if err != nil {
		http.Error(w, "Error encoding product", http.StatusInternalServerError)
		return
	}
}
func main() {
	// Create a new router
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/products",getProducts)
	mux.HandleFunc("/create-product", createProductHandler)
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