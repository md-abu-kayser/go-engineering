package main

import "fmt"

func summarizeExportedIdentifiers() (string, int) {
	topic := "Exported Identifiers"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeExportedIdentifiers()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
