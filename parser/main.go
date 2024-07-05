/* silver-data-parser-go

Examples:
go run . -task=test
*/

package main

import (
	"flag"
	"log"
	"silver-data/parser-go/test"
)

func main() {
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

type options struct {
	task      string
	directory string
}

func getOptions() options {
	var result options
	taskPtr := flag.String("task", "none", "The task for the parser to perform")
	directoryPtr := flag.String("directory", "./test/", "Directory for tests")
	flag.Parse()

	result.task = *taskPtr
	result.directory = *directoryPtr

	return result
}
