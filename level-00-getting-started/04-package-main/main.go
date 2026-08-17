package main

import "fmt"

func PackageMain() string {
	const topic = "Package Main"
	return topic
}

func main() {
	fmt.Println(PackageMain())
}
