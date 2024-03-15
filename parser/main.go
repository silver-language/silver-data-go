package main

import (
	"fmt"
	"os"
)

func main() {

	// read file

	filename := "test.txt"

	// show file stats
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file.Stat())

	// read the file
	content, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Err")
	}
	fmt.Println(string(content))

	defer file.Close()

}
