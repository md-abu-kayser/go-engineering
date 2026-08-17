package main

import "fmt"

func summarizeCsvImportService() (string, int) {
	topic := "Csv Import Service"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCsvImportService()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
