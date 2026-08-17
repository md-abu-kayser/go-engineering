package main

import "fmt"

func SignalNotify() string {
	const topic = "Signal Notify"
	return topic
}

func main() {
	fmt.Println(SignalNotify())
}
