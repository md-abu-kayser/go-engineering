package main

import "fmt"

func ExecWithTimeout() string {
	const topic = "Exec With Timeout"
	return topic
}

func main() {
	fmt.Println(ExecWithTimeout())
}
