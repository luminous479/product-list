package main

import (

	"fmt"

	"github.com/luminous479/product-list/utils"
)

//"github.com/luminous479/product-list/cmd"

func main() {
	//cmd.Serve()

	// hmac-sha-256


token , err := utils.CreateJwt("my-secret",utils.PayLoad{
	Sub: 28,
	Name: "Akhi",
	Email: "akhi@example.com",
	IsShopOwner: true,
})	
if err != nil {
	fmt.Println("Error creating JWT:", err)
	return

}
fmt.Println("Generated JWT:", token)
}