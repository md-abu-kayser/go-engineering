package main

import "fmt"

func MultipleReturnValues() string {
	const topic = "Multiple Return Values"
	return topic
}

func main() {
	fmt.Println(MultipleReturnValues())
}
