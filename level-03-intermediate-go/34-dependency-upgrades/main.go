package main

import "fmt"

func summarizeDependencyUpgrades() (string, int) {
	topic := "Dependency Upgrades"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeDependencyUpgrades()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
