package repository

import (
	"database/sql"
	"fmt"
	"gin-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id , name , price FROM produto"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var prodcutObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&prodcutObj.ID,
			&prodcutObj.Name,
			&prodcutObj.Price,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Product{}, err
		}

		productList = append(productList, prodcutObj)

	}

	rows.Close()

	return productList, nil

}

func (pr *ProductRepository) CreatProduct(product model.Product) (int, error) {

	var id int
	query, err := pr.connection.Prepare("INSERT INTO produto " +
		"(name,price) VALUES ($1,$2) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil

}

func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {

	query, err := pr.connection.Prepare("SELECT * FROM produto WHERE id  = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var produto model.Product

	err = query.QueryRow(id).Scan(
		&produto.ID,
		&produto.Name,
		&produto.Price,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	query.Close()

	return &produto, nil

}
