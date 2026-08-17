package main

import "fmt"

func summarizeHtmlTemplate() (string, int) {
	topic := "Html Template"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHtmlTemplate()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
