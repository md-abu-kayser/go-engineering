package main

import "fmt"

func summarizeDarkLaunch() (string, int) {
	topic := "Dark Launch"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDarkLaunch()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
