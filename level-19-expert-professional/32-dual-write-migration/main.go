package main

import "fmt"

func summarizeDualWriteMigration() (string, int) {
	topic := "Dual Write Migration"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDualWriteMigration()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
