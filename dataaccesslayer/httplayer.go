package dataaccesslayer

import (
	"github.com/labstack/echo/v4"

	"productapi/services"
)

func StartHttpServer() {
	server := echo.New()

	server.GET("/getallproducts", services.GetAllProducts)

	server.POST("/addproduct", services.AddProduct)

	server.Logger.Fatal(server.Start(":1234"))
}
