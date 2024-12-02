package utils

import (
	"os"
	"log"
)

func ReadFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}