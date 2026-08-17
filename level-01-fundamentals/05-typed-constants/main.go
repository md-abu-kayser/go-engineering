package main

import "fmt"

func TypedConstants() string {
	const topic = "Typed Constants"
	return topic
}

func main() {
	fmt.Println(TypedConstants())
}
