package main

import "fmt"

func StubPattern() string {
	const topic = "Stub Pattern"
	return topic
}

func main() {
	fmt.Println(StubPattern())
}
