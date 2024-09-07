package services

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"productapi/databaselayer"
)

func GetAllProducts(client echo.Context) error {

	var sqlDb = databaselayer.GetConnection()

	var iproductService databaselayer.IProductService = &databaselayer.ProductService{SqlDb: sqlDb}

	var products = iproductService.GetAllProducts()

	return client.JSON(http.StatusOK, products)
}
