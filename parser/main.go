/* silver-data-parser-go

Examples:
go run . -task=test
*/

package main

import (
	"log"
	"silver-data/parser-go/test"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Module: silver-data-parser-go")

	options := getOptions()

	log.Println("Task:", options.task)
	switch options.task {
	case "test":
		test.TestDirectory(options.directory)
	case "generate":
		test.GenerateDirectoryOutput(options.directory)
	default:
		{
			log.Println("No matching task")
			log.Println("Tasks: test, generate")
		}
	}
}
