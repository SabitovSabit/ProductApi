package services

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"productapi/databaselayer"
	"productapi/models"
)

func GetAllProducts(client echo.Context) error {

	var sqlDb = databaselayer.GetConnection()

	var iproductService databaselayer.IProductService = &databaselayer.ProductService{SqlDb: sqlDb}

	var products = iproductService.GetAllProducts()

	return client.JSON(http.StatusOK, products)
}

func AddProduct(client echo.Context) error {

	product := new(models.Product)

	err := client.Bind(product)

	if err != nil {
		return client.JSON(http.StatusBadRequest, "Invalid input")
	}

	var sqlDb = databaselayer.GetConnection()

	var iproductService databaselayer.IProductService = &databaselayer.ProductService{SqlDb: sqlDb}

	var resp = iproductService.AddProduct(*product)

	return client.JSON(http.StatusOK, resp)
}

func DeleteProduct(client echo.Context) error {

	param := client.QueryParam("id")

	Id, err := strconv.Atoi(param)

	if err != nil {
		return client.JSON(http.StatusBadRequest, "Invalid input")
	}

	var sqlDb = databaselayer.GetConnection()

	var iproductService databaselayer.IProductService = &databaselayer.ProductService{SqlDb: sqlDb}

	var resp = iproductService.DeleteProduct(Id)

	return client.JSON(http.StatusOK, resp)
}

func UpdateProduct(client echo.Context) error {

	product := new(models.Product)

	err := client.Bind(product)

	if err != nil {
		return client.JSON(http.StatusBadRequest, "Invalid input")
	}

	var sqlDb = databaselayer.GetConnection()

	var iproductService databaselayer.IProductService = &databaselayer.ProductService{SqlDb: sqlDb}

	var resp = iproductService.UpdateProduct(*product)

	return client.JSON(http.StatusOK, resp)
}
