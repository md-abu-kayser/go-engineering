package main

import "fmt"

func summarizeTlsMinVersion() (string, int) {
	topic := "Tls Min Version"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTlsMinVersion()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
