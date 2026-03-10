package main

import (
	"fmt"
	"net/http"
) 
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func main() {
	// Create a new router
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	// Start the server
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}