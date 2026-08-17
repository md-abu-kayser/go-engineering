package main

import "fmt"

func StaticFiles() string {
	const topic = "Static Files"
	return topic
}

func main() {
	fmt.Println(StaticFiles())
}
