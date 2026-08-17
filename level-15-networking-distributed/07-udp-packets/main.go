package main

import "fmt"

func summarizeUdpPackets() (string, int) {
	topic := "Udp Packets"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeUdpPackets()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
