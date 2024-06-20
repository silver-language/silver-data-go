package file

import (
	"fmt"
	"os"
)

func Read(filepath string) string {
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error opening file: %v \n", err)
	}
	defer file.Close()

	fmt.Printf("File opened: %v \n", filepath)

	// Read file
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v \n", err)
	}
	fmt.Printf("File read: %v \n", filepath)

	return string(content)
}

func Write(filepath string, content string) {
	//fmt.Printf("filepath: %v \n", filepath)
	err := os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		fmt.Printf("Error writing file: %v \n", err)
	}
	fmt.Printf("File written: %v \n", filepath)
}
