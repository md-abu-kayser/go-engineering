package main

import "fmt"

func summarizeOauthState() (string, int) {
	topic := "Oauth State"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeOauthState()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
