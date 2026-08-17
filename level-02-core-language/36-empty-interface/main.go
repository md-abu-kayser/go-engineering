package main

import "fmt"

func EmptyInterface() string {
	const topic = "Empty Interface"
	return topic
}

func main() {
	fmt.Println(EmptyInterface())
}
