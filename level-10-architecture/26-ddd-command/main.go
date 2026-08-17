package main

import "fmt"

func summarizeDddCommand() (string, int) {
	topic := "Ddd Command"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDddCommand()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
