package main

import (
	"fmt"
	"os"
	"silver-data/parser-go/test"
)

func main() {
	fmt.Println("Module: silver-data-parser-go")

	if len(os.Args) > 1 {
		task := os.Args[1]
		fmt.Println("Task:", task)
		switch task {
		case "test":
			test.TestDirectory("./test/")
		case "generate":
			test.GenerateDirectoryOutput("./test/")
		default:
			fmt.Println("No matching task")
		}
	} else {
		fmt.Println("Tasks: test, generate")
	}
}
