// File: main.go

package main

import (
	"ekeberg.com/go-api-sql-gcp-products/db"
	"ekeberg.com/go-api-sql-gcp-products/handlers"
	"ekeberg.com/go-api-sql-gcp-products/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	// SQLite connection
	db.InitDB()

	// Start Gin Router
	r := gin.Default()

	// Serve static files for the favicon
	r.StaticFile("/favicon.ico", "./assets/favicon.ico")

	// API v1
	v1 := r.Group("/api/v1")
	{
		// Users (no authentication required)
		v1.POST("users/signup", handlers.SignUp) // POST http://localhost:8080/api/v1/users/signup
		v1.POST("users/login", handlers.Login)   // POST http://localhost:8080/api/v1/users/login

		// Products (authentication required as human or service)
		authenticatedHumanOrService := v1.Group("/")
		authenticatedHumanOrService.Use(middlewares.Authenticate)
		{
			authenticatedHumanOrService.GET("products", handlers.GetProducts)        // GET http://localhost:8080/api/v1/products
			authenticatedHumanOrService.GET("products/:id", handlers.GetProductById) // GET http://localhost:8080/api/v1/products/1
		}

		// Products (authentication required as human)
		authenticatedHumanOnly := v1.Group("/")
		authenticatedHumanOnly.Use(middlewares.Authenticate)
		{
			authenticatedHumanOnly.POST("products", handlers.AddProduct)          // POST http://localhost:8080/api/v1/products
			authenticatedHumanOnly.PUT("products/:id", handlers.UpdateProduct)    // PUT http://localhost:8080/api/v1/products/17
			authenticatedHumanOnly.DELETE("products/:id", handlers.DeleteProduct) // DELETE http://localhost:8080/api/v1/products/17
			authenticatedHumanOnly.OPTIONS("products", handlers.Options)          // OPTIONS http://localhost:8080/api/v1/products
		}
	}

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	r.Run()
}
