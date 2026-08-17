package main

import "fmt"

func summarizeRwmutex() (string, int) {
	topic := "Rwmutex"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRwmutex()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
