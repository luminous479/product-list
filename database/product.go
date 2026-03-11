
package database

var products  []Product
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func AddProduct(product Product)  Product{
	product.ID = len(products) + 1
	products = append(products, product)
	return product
}

func GetProducts() []Product {
	return products
}

func GetProductByID(id int) Product {
	for _, product := range products {
		if product.ID == id {
			return product
		}
	}
	return Product{}
}

func UpdateProduct(id int, updatedProduct Product) Product {
	updatedProduct.ID = id
	for i, product := range products {
		if product.ID == id {
			products[i] = updatedProduct
			return updatedProduct
		}
	}
	return Product{}
}
func DeleteProduct(id int) Product {	
	for i, product := range products {
		if product.ID == id {
			products = append(products[:i], products[i+1:]...)	
			return product
		}
	}
	return Product{}
}


func init() {
	products = []Product{
		{ID: 1, Name: "Product 1", Price: 10.99},
		{ID: 2, Name: "Product 2", Price: 19.99},
		{ID: 3, Name: "Product 3", Price: 5.49},
	}
	
}