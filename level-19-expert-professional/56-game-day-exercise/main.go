package main

import "fmt"

func summarizeGameDayExercise() (string, int) {
	topic := "Game Day Exercise"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeGameDayExercise()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
