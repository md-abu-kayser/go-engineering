package main

import "fmt"

func CgoPointerRules() string {
	const topic = "Cgo Pointer Rules"
	return topic
}

func main() {
	fmt.Println(CgoPointerRules())
}
