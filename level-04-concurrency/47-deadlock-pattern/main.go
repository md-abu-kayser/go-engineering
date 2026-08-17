package main

import "fmt"

func DeadlockPattern() string {
	const topic = "Deadlock Pattern"
	return topic
}

func main() {
	fmt.Println(DeadlockPattern())
}
