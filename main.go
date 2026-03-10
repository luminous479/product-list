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

func corsHandler(w http.ResponseWriter){
	
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

	}

func optionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func sendData(w http.ResponseWriter, data interface{}, statusCode int) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		http.Error(w, "Error encoding data", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	corsHandler(w)
	optionsHandler(w, r)
	
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	sendData(w, products, http.StatusOK)

}
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	corsHandler(w)
	optionsHandler(w, r)

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
	sendData(w, newProduct, http.StatusCreated)
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