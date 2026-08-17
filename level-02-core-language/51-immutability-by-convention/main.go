package main

import "fmt"

func ImmutabilityByConvention() string {
	const topic = "Immutability By Convention"
	return topic
}

func main() {
	fmt.Println(ImmutabilityByConvention())
}
