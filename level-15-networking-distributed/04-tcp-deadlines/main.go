package main

import "fmt"

func summarizeTcpDeadlines() (string, int) {
	topic := "Tcp Deadlines"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTcpDeadlines()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
