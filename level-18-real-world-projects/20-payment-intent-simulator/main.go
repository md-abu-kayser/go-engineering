package main

import "fmt"

func summarizePaymentIntentSimulator() (string, int) {
	topic := "Payment Intent Simulator"
	return topic, len(topic)
}

func main() {
	topic, length := summarizePaymentIntentSimulator()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
