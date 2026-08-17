package main

import "fmt"

func summarizeInboxTable() (string, int) {
	topic := "Inbox Table"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeInboxTable()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
