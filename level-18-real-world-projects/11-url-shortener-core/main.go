package main

import "fmt"

func summarizeUrlShortenerCore() (string, int) {
	topic := "Url Shortener Core"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUrlShortenerCore()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
