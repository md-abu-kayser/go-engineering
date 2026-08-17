package main

import "fmt"

func GcMarkPhase() string {
	const topic = "Gc Mark Phase"
	return topic
}

func main() {
	fmt.Println(GcMarkPhase())
}
