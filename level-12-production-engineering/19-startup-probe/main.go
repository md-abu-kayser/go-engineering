package main

import "fmt"

func StartupProbe() string {
	const topic = "Startup Probe"
	return topic
}

func main() {
	fmt.Println(StartupProbe())
}
