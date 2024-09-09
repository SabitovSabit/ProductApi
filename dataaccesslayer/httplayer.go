package dataaccesslayer

import (
	"github.com/labstack/echo/v4"

	"productapi/services"
)

func StartHttpServer() {
	server := echo.New()

	server.GET("/getallproducts", services.GetAllProducts)

	server.POST("/addproduct", services.AddProduct)

	server.DELETE("/deleteproduct", services.DeleteProduct)

	server.PUT("/updateproduct", services.UpdateProduct)

	server.Logger.Fatal(server.Start(":1234"))
}
