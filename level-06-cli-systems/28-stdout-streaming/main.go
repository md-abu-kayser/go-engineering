package main

import "fmt"

func StdoutStreaming() string {
	const topic = "Stdout Streaming"
	return topic
}

func main() {
	fmt.Println(StdoutStreaming())
}
