package main

import "fmt"

func RequestHeaders() string {
	const topic = "Request Headers"
	return topic
}

func main() {
	fmt.Println(RequestHeaders())
}
