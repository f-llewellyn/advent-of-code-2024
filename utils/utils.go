package utils

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
	}

	return data
}
