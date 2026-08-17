package main

import "fmt"

func CacheAside() string {
	const topic = "Cache Aside"
	return topic
}

func main() {
	fmt.Println(CacheAside())
}
