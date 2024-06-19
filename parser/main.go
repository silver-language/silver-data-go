/* silver-data-parser-go

Examples:
go run . -task=test
*/

package main

import (
	"flag"
	"fmt"
	"silver-data/parser-go/test"
)

func main() {
	fmt.Println("Module: silver-data-parser-go")

	taskPtr := flag.String("task", "none", "The task for the parser to perform")
	directoryPtr := flag.String("directory", "./test/", "Directory for tests")
	flag.Parse()

	task := *taskPtr // os.Args[1]
	fmt.Println("Task:", task)
	switch task {
	case "test":
		test.TestDirectory(*directoryPtr)
	case "generate":
		test.GenerateDirectoryOutput(*directoryPtr)
	default:
		{
			fmt.Println("No matching task")
			fmt.Println("Tasks: test, generate")
		}
	}
}
