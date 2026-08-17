package main

import "fmt"

func BufferReuse() string {
	const topic = "Buffer Reuse"
	return topic
}

func main() {
	fmt.Println(BufferReuse())
}
