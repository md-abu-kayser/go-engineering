package main

import "fmt"

func summarizeGoModTidy() (string, int) {
	topic := "Go Mod Tidy"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGoModTidy()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
