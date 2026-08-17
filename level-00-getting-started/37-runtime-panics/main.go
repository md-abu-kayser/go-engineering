package main

import "fmt"

func RuntimePanics() string {
	const topic = "Runtime Panics"
	return topic
}

func main() {
	fmt.Println(RuntimePanics())
}
