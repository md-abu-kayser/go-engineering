package main

import "fmt"

func BoundedConcurrency() string {
	const topic = "Bounded Concurrency"
	return topic
}

func main() {
	fmt.Println(BoundedConcurrency())
}
