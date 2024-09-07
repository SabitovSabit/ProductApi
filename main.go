package main

import (
	"fmt"
	"productapi/databaselayer"
	"productapi/services"
)

func main() {
	fmt.Println("Hello!")

	var sqlDb = databaselayer.GetConnection()

	var iproductService services.IProductService = &services.ProductService{SqlDb: sqlDb}

	var products = iproductService.GetAllProducts()

	for _, pro := range products {
		fmt.Println(pro.Name)
	}

}
