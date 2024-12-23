package services

import (
	"database/sql"
	"vending_machine/dtos"
)

type ProductService interface {
	Service[dtos.Product]
	UpdateProduct(id int, quantity int) error
}

type productService struct {
	Service[dtos.Product]
	mysqlDB *sql.DB
}

func NewProductService(mySqlConn *sql.DB) ProductService {
	return &productService{
		Service: NewService[dtos.Product](mySqlConn, dtos.DBProduct),
		mysqlDB: mySqlConn,
	}
}

func (c *productService) UpdateProduct(id int, quantity int) error {
	_, err := c.mysqlDB.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", quantity, id)
	return err
}

