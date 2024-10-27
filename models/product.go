// File: models/product.go

package models

import (
	"database/sql"
	"ekeberg.com/go-api-sql-gcp-products/db"
	"fmt"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Ean         string `json:"ean"`
	PriceOut    string `json:"price_out"`
}

func GetProducts(count int) ([]Product, error) {
	// Initialize the base query
	query := "SELECT id, name, description, ean, price_out FROM products"

	// Modify the query to include a LIMIT if count is greater than zero
	if count > 0 {
		query += " LIMIT " + strconv.Itoa(count)
	} else if count == 0 {
		// If count is 0, default it to 10000
		count = 10000
		query += " LIMIT " + strconv.Itoa(count)
	}

	// Execute the query
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize an empty slice to store products
	products := make([]Product, 0)

	// Iterate through the rows and scan into Product structs
	for rows.Next() {
		singleProduct := Product{}
		err = rows.Scan(&singleProduct.Id, &singleProduct.Name, &singleProduct.Description, &singleProduct.Ean, &singleProduct.PriceOut)
		if err != nil {
			return nil, err
		}
		products = append(products, singleProduct)
	}

	// Check for any errors encountered during row iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func GetProductById(id string) (Product, error) {

	stmt, err := db.DB.Prepare("SELECT id, name, description, ean, price_out FROM products WHERE id = ?")

	if err != nil {
		return Product{}, err
	}

	product := Product{}

	sqlErr := stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Description, &product.Ean, &product.PriceOut)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return Product{}, nil
		}
		return Product{}, sqlErr
	}
	return product, nil
}

func AddProduct(newProduct Product) (bool, error) {

	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("INSERT INTO products (name, description, ean, price_out) VALUES (?, ?, ?, ?)")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(newProduct.Name, newProduct.Description, newProduct.Ean, newProduct.PriceOut)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func UpdateProduct(ourProduct Product, id int) (bool, error) {

	tx, err := db.DB.Begin()
	if err != nil {
		return false, err
	}

	stmt, err := tx.Prepare("UPDATE products SET name = ?, description = ?, ean = ?, price_out = ? WHERE id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ourProduct.Name, ourProduct.Description, ourProduct.Ean, ourProduct.PriceOut, id)

	fmt.Print("UPDATE products SET name='{ourProduct.Name}', description='{ourProduct.Description}', ean='{ourProduct.Ean}', price_out={ourProduct.PriceOut} WHERE id={id}")

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}

func DeleteProduct(personId int) (bool, error) {

	tx, err := db.DB.Begin()

	if err != nil {
		return false, err
	}

	stmt, err := db.DB.Prepare("DELETE FROM products where id = ?")

	if err != nil {
		return false, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(personId)

	if err != nil {
		return false, err
	}

	tx.Commit()

	return true, nil
}
