package main

import "fmt"

func FakeRandomness() string {
	const topic = "Fake Randomness"
	return topic
}

func main() {
	fmt.Println(FakeRandomness())
}
