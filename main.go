// File: main.go

package main

import (
	"ekeberg.com/go-api-sql-gcp-products/handlers"
	"ekeberg.com/go-api-sql-gcp-products/models"
	"ekeberg.com/go-api-sql-gcp-products/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// SQLite connection
	err := models.ConnectDatabase()
	utils.CheckErr(err)

	// Start Gin Router
	r := gin.Default()

	// Serve static files for the favicon
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("product", handlers.GetProducts) // http://localhost:8080/api/v1/product
		v1.GET("product/:id", handlers.GetProductById)
		v1.POST("product", handlers.AddProduct)
		v1.PUT("product/:id", handlers.UpdateProduct)
		v1.DELETE("product/:id", handlers.DeleteProduct)
		v1.OPTIONS("product", handlers.Options)
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}
