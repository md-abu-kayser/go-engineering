package main

import "fmt"

func summarizeDeferLifoOrder() (string, int) {
	topic := "Defer Lifo Order"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDeferLifoOrder()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
