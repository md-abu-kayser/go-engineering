package main

import "fmt"

func summarizeExecStatement() (string, int) {
	topic := "Exec Statement"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeExecStatement()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
