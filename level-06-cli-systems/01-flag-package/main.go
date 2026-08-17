package main

import "fmt"

func FlagPackage() string {
	const topic = "Flag Package"
	return topic
}

func main() {
	fmt.Println(FlagPackage())
}
