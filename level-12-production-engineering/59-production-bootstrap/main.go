package main

import "fmt"

func summarizeProductionBootstrap() (string, int) {
	topic := "Production Bootstrap"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeProductionBootstrap()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
