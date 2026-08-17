package main

import "fmt"

func summarizeMutexProfilerDeep() (string, int) {
	topic := "Mutex Profiler Deep"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeMutexProfilerDeep()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
