package main

import "fmt"

func summarizeRegexpBasic() (string, int) {
	topic := "Regexp Basic"
	return topic, len(topic)
}

func main() {
	topic, length := summarizeRegexpBasic()
	fmt.Printf("%s (%d chars)\n", topic, length)
}
