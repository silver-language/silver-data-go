/* silver-data-parser-go

Examples:
go run . -task=test
go run . -task=generate
go run . -task=generate -directory=test/simple/
go run . -task=generateFile -file=test/simple/array.agd
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
		test.TestDirectory(options.directory) // currently doesn't do much
	case "generate":
		test.GenerateDirectoryOutput(options.directory)
	case "generateFile":
		test.GenerateFileOutput(options.file)
	default:
		{
			log.Println("No matching task")
			log.Println("Tasks: test, generate")
		}
	}
}
