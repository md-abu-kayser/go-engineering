package main

import "fmt"

func MessageQueueProducer() string {
	const topic = "Message Queue Producer"
	return topic
}

func main() {
	fmt.Println(MessageQueueProducer())
}
