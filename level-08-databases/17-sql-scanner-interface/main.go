package main

import "fmt"

func summarizeSqlScannerInterface() (string, int) {
	topic := "Sql Scanner Interface"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeSqlScannerInterface()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
