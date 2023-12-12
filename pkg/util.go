package pkg

import (
	"app/internal/domain"
	"encoding/json"
	"io"
	"log"
	"os"
)

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
