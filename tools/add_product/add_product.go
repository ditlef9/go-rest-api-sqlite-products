// File: tools/add_product/add_product.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Product struct
type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Ean         string `json:"ean"`
	PriceOut    string `json:"price_out"`
}

func main() {
	var product Product

	// Prompt user for input
	fmt.Print("Enter name: ")
	fmt.Scanln(&product.Name)

	fmt.Print("Enter description: ")
	fmt.Scanln(&product.Description)

	fmt.Print("Enter EAN: ")
	fmt.Scanln(&product.Ean)

	fmt.Print("Enter price out: ")
	fmt.Scanln(&product.PriceOut)

	// Marshal the person struct into JSON
	jsonData, err := json.Marshal(product)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		os.Exit(1)
	}

	// Send POST request
	response, err := http.Post("http://localhost:8080/api/v1/product", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode == http.StatusOK {
		fmt.Println("Data successfully sent.")
	} else {
		fmt.Printf("Failed to send data. Status code: %d\n", response.StatusCode)
	}
}
