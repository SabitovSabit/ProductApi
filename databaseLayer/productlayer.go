package databaselayer

import (
	"context"
	"database/sql"
	"productapi/models"

	"github.com/labstack/gommon/log"
)

type IProductService interface {
	GetAllProducts() []models.Product
}

type ProductService struct {
	SqlDb *sql.DB
}

func (productService *ProductService) GetAllProducts() []models.Product {
	rows, err := productService.SqlDb.QueryContext(context.TODO(), "select Id,Name,Price,Count from Product")

	if err != nil {
		log.Error("unable select all products %v\n", err)
	}

	defer rows.Close()
	defer productService.SqlDb.Close()

	var allProducts []models.Product

	for rows.Next() {
		var produtc models.Product

		rows.Scan(&produtc.Id, &produtc.Name, &produtc.Price, &produtc.Count)

		allProducts = append(allProducts, produtc)
	}

	return allProducts
}
