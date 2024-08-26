package main

import "flag"

type options struct {
	task      string
	directory string
	file      string
}

func getOptions() options {
	var result options
	taskPtr := flag.String("task", "none", "The task for the parser to perform")
	directoryPtr := flag.String("directory", "./test/", "Directory for tests")
	filePtr := flag.String("file", "./test/simple/array.agd", "File to test")
	flag.Parse()

	result.task = *taskPtr
	result.directory = *directoryPtr
	result.file = *filePtr

	return result
}
