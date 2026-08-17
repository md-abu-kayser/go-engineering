package main

import "fmt"

func summarizeRpcServer() (string, int) {
	topic := "Rpc Server"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRpcServer()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
