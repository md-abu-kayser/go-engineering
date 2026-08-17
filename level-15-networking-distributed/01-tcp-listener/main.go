package main

import "fmt"

func summarizeTcpListener() (string, int) {
	topic := "Tcp Listener"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTcpListener()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
