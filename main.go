package main

import (
	"encoding/base64"
	"fmt"
)

//"github.com/luminous479/product-list/cmd"

func main() {
	//cmd.Serve()
	s := "abc"
	byteArr := []byte(s)

	enc := base64.URLEncoding
	enc = enc.WithPadding(base64.NoPadding)
	b64Str := enc.EncodeToString(byteArr)
    fmt.Println("Byte array : ",byteArr)
	fmt.Println("Base 64 : ",b64Str)


}
