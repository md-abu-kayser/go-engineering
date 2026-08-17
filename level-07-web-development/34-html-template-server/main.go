package main

import "fmt"

func summarizeHtmlTemplateServer() (string, int) {
	topic := "Html Template Server"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeHtmlTemplateServer()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
