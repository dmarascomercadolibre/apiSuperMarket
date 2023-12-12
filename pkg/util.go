package pkg

import (
	"app/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
)

// FullfilDBProduct reads a JSON file from the given path and returns a slice of domain.Product.
// It opens the file, reads its contents, and unmarshals the JSON data into a slice of domain.Product.
// If any error occurs during the process, it will log the error and terminate the program.
func FullfilDBProduct(path string) []domain.Product {
	data, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	dataRead, err := io.ReadAll(data)
	if err != nil {
		log.Fatal(err)
	}
	slice := []domain.Product{}
	json.Unmarshal(dataRead, &slice)

	return slice
}
