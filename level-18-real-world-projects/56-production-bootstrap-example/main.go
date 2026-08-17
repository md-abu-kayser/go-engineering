package main

import "fmt"

func summarizeProductionBootstrapExample() (string, int) {
	topic := "Production Bootstrap Example"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeProductionBootstrapExample()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
