package main

import "fmt"

func summarizeZipSlipDefense() (string, int) {
	topic := "Zip Slip Defense"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeZipSlipDefense()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
