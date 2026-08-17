package main

import "fmt"

func summarizeMigrationLock() (string, int) {
	topic := "Migration Lock"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMigrationLock()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
