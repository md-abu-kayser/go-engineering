package main

import "fmt"

func summarizeSortSearch() (string, int) {
	topic := "Sort Search"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSortSearch()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
