package main

import (
	"crypto/sha256"
	
	"fmt"
)

//"github.com/luminous479/product-list/cmd"

func main() {
	//cmd.Serve()

	// sha - 256
	data := []byte("Hello World!")
	hash := sha256.Sum256(data)
	fmt.Println("Hash after SHA-256 : ", hash)
	

}
