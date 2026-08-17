package main

import "fmt"

func summarizeAdapterCli() (string, int) {
	topic := "Adapter Cli"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeAdapterCli()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
