package main

import "fmt"

func DeadLetterQueue() string {
	const topic = "Dead Letter Queue"
	return topic
}

func main() {
	fmt.Println(DeadLetterQueue())
}
