package main

import "fmt"

func TraceSpan() string {
	const topic = "Trace Span"
	return topic
}

func main() {
	fmt.Println(TraceSpan())
}
