package main

import "fmt"

func AtomicAdd() string {
	const topic = "Atomic Add"
	return topic
}

func main() {
	fmt.Println(AtomicAdd())
}
