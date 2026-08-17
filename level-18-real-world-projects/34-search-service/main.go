package main

import "fmt"

func SearchService() string {
	const topic = "Search Service"
	return topic
}

func main() {
	fmt.Println(SearchService())
}
