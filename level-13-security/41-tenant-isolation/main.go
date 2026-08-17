package main

import "fmt"

func summarizeTenantIsolation() (string, int) {
	topic := "Tenant Isolation"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeTenantIsolation()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
