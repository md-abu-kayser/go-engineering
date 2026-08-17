package main

import "fmt"

func IoCopy() string {
	const topic = "Io Copy"
	return topic
}

func main() {
	fmt.Println(IoCopy())
}
