package dataaccesslayer

import (
	"github.com/labstack/echo/v4"

	"productapi/services"
)

func StartHttpServer() {
	server := echo.New()

	server.GET("/GetAllProducts", services.GetAllProducts)

	server.Logger.Fatal(server.Start(":1234"))
}
