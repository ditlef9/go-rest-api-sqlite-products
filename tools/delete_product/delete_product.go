// File: tools/delete_product/delete_product.go
package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	var id string

	// Prompt user for ID
	fmt.Print("Enter ID to delete: ")
	fmt.Scanln(&id)

	// Send DELETE request to delete the person
	url := "http://localhost:8080/api/v1/product/" + id
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Error creating DELETE request:", err)
		os.Exit(1)
	}

	// Execute the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending DELETE request:", err)
		os.Exit(1)
	}
	defer response.Body.Close()

	// Check response status
	if response.StatusCode == http.StatusOK {
		fmt.Println("Successfully deleted.")
	} else {
		fmt.Printf("Failed to delete. Status code: %d\n", response.StatusCode)
	}
}
