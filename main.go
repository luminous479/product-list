package main

import (
	"crypto/hmac"
	"crypto/sha256"

	"fmt"
)

//"github.com/luminous479/product-list/cmd"

func main() {
	//cmd.Serve()

	// hmac-sha-256

	secret := []byte("my-secret")
	msg := []byte("hello world!")

	h := hmac.New(sha256.New,secret)
	h.Write(msg)

	text := h.Sum(nil)
	fmt.Println(text)

	

}
