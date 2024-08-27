package util

import (
	"log"
	"os"
)

func FileRead(filepath string) string {
	// Open file
	file, err := os.Open(filepath)
	if err != nil {
		log.Printf("Error opening file: %v \n", err)
	}
	defer file.Close()

	log.Printf("File opened: %v \n", filepath)

	// Read file
	content, err := os.ReadFile(filepath)
	if err != nil {
		log.Printf("Error reading file: %v \n", err)
	}
	log.Printf("File read: %v \n", filepath)

	return string(content)
}

func FileWrite(filepath string, content string) {
	//log.Printf("filepath: %v \n", filepath)
	err := os.WriteFile(filepath, []byte(content), 0644)
	if err != nil {
		log.Printf("Error writing file: %v \n", err)
	}
	log.Printf("File written: %v \n", filepath)
}
