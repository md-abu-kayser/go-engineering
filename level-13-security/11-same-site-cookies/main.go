package main

import "fmt"

func summarizeSameSiteCookies() (string, int) {
	topic := "Same Site Cookies"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSameSiteCookies()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
