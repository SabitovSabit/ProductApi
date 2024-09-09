package databaselayer

import (
	"context"
	"database/sql"
	"productapi/models"

	"github.com/labstack/gommon/log"
)

type IProductService interface {
	GetAllProducts() []models.Product

	AddProduct(product models.Product) models.Product

	UpdateProduct(product models.Product) models.Product

	DeleteProduct(id int) string
}

type ProductService struct {
	SqlDb *sql.DB
}

func (productService *ProductService) GetAllProducts() []models.Product {
	rows, err := productService.SqlDb.QueryContext(context.TODO(), "select Id,Name,Price,Count from Product")

	if err != nil {
		log.Errorf("unable select all products %v\n", err)
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

func (productService *ProductService) AddProduct(product models.Product) models.Product {
	defer productService.SqlDb.Close()

	command := "insert into Product (name,price,count) values (@p1,@p2,@p3);SELECT SCOPE_IDENTITY();"

	var newId int
	err := productService.SqlDb.QueryRow(command, product.Name, product.Price, product.Count).Scan(&newId)

	if err != nil {
		log.Errorf("unable inser one product %v\n", err)
	}

	product.Id = newId

	return product
}

func (productService *ProductService) DeleteProduct(id int) string {

	defer productService.SqlDb.Close()

	command := "delete product where Id=@p1"

	result, err := productService.SqlDb.ExecContext(context.TODO(), command, id)

	if err != nil {
		log.Errorf("unable delete one product %d %v", id, err)
		return "Error"
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		log.Errorf("error fetching product rows %v", err)
		return "error fetching product"

	}

	if rowsAffected == 0 {
		log.Errorf("no product found with id %d", id)
		return "no product found"
	}

	return "Deleted"
}

func (productService *ProductService) UpdateProduct(product models.Product) models.Product {
	defer productService.SqlDb.Close()

	command := "update Product set name = @p1,price=@p2,count=@p3 where Id=@p4;"

	_, err := productService.SqlDb.Exec(command, product.Name, product.Price, product.Count, product.Id)

	if err != nil {
		log.Errorf("unable update one product %v", err)
	}

	return product
}
