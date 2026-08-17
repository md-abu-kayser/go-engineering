package main

import "fmt"

func summarizeSessionCookieDesign() (string, int) {
	topic := "Session Cookie Design"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSessionCookieDesign()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
