package main

import "fmt"

func DataRaceExample() string {
	const topic = "Data Race Example"
	return topic
}

func main() {
	fmt.Println(DataRaceExample())
}
