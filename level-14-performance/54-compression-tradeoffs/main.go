package main

import "fmt"

func summarizeCompressionTradeoffs() (string, int) {
	topic := "Compression Tradeoffs"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeCompressionTradeoffs()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
