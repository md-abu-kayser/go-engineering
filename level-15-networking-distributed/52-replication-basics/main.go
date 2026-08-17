package main

import "fmt"

func summarizeReplicationBasics() (string, int) {
	topic := "Replication Basics"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeReplicationBasics()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
