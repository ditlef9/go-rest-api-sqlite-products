// File: handlers/product_handlers.go

package handlers

import (
	"ekeberg.com/go-api-sql-gcp-products/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products, err := models.GetProducts(-1)
	if err != nil {
		log.Fatal(err)
	}

	if products == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": products})
	}
}

func GetProductById(c *gin.Context) {
	id := c.Param("id")
	product, err := models.GetProductById(id)
	if err != nil {
		log.Fatal(err)
	}

	if product.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Records Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": product})
	}
}

func AddProduct(c *gin.Context) {
	var json models.Product

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	success, err := models.AddProduct(json)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func UpdateProduct(c *gin.Context) {
	var json models.Product

	// Debugging print to show empty struct before binding
	fmt.Printf("products_h.go::UpdateProduct()::Initial empty Payload struct: %+v\n", json)

	// Attempt to bind the incoming JSON request to the Product struct
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Printf("products_h.go::UpdateProduct()::Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debugging print to show the parsed JSON payload
	fmt.Printf("products_h.go::UpdateProduct()::Parsed JSON Payload: %+v\n", json)

	// Get the product ID from the request path
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Printf("products_h.go::UpdateProduct()::Error converting ID: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Debugging print to show the product ID to be updated
	fmt.Printf("products_h.go::UpdateProduct()::Updating product with ID: %d\n", productId)

	// Attempt to update the product in the database
	success, err := models.UpdateProduct(json, productId)
	if err != nil {
		fmt.Printf("products_h.go::UpdateProduct()::Error updating product: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the update was successful
	if success {
		fmt.Printf("products_h.go::UpdateProduct()::Product with ID %d successfully updated\n", productId)
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		fmt.Printf("products_h.go::UpdateProduct()::Failed to update product with ID %d\n", productId)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func DeleteProduct(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}

	success, err := models.DeleteProduct(productId)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func Options(c *gin.Context) {
	ourOptions := "HTTP/1.1 200 OK\n" +
		"Allow: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Origin: http://locahost:8080\n" +
		"Access-Control-Allow-Methods: GET,POST,PUT,DELETE,OPTIONS\n" +
		"Access-Control-Allow-Headers: Content-Type\n"

	c.String(200, ourOptions)
}
