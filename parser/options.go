package main

import "flag"

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
