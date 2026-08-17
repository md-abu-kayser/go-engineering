package main

import "fmt"

func summarizeSchemaVersioning() (string, int) {
	topic := "Schema Versioning"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSchemaVersioning()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
