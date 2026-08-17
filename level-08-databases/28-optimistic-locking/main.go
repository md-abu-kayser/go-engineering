package main

import "fmt"

func OptimisticLocking() string {
	const topic = "Optimistic Locking"
	return topic
}

func main() {
	fmt.Println(OptimisticLocking())
}
