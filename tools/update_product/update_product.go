// File: tools/update_product.go
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
	var id string

	// Prompt user for ID
	fmt.Print("Enter ID to update: ")
	fmt.Scanln(&id)

	// Prompt user for new input data
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

	// Send PUT request to update the person
	url := "http://localhost:8080/api/v1/product/" + id
	fmt.Printf("JSON Payload: %s\n", jsonData)
	request, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating PUT request:", err)
		os.Exit(1)
	}
	request.Header.Set("Content-Type", "application/json")

	// Execute the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending PUT request:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode == http.StatusOK {
		fmt.Println("Data successfully updated.")
	} else {
		fmt.Printf("Failed to update data. Status code: %d\n", response.StatusCode)
	}
}
