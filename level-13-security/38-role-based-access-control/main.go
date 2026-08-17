package main

import "fmt"

func summarizeRoleBasedAccessControl() (string, int) {
	topic := "Role Based Access Control"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRoleBasedAccessControl()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
