package main

import "fmt"

func TraceTask() string {
	const topic = "Trace Task"
	return topic
}

func main() {
	fmt.Println(TraceTask())
}
