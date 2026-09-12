// Lesson 57: Example Documentation
//
// Goal: Keep a package boundary narrow by depending on the behavior a
// consumer needs rather than a concrete infrastructure type.
package main

import "fmt"

const lesson = "Example Documentation"

type notifier interface {
	Send(string) string
}

type consoleNotifier struct{}

func (consoleNotifier) Send(message string) string {
	return "sent: " + message
}

func notify(n notifier, message string) string {
	return n.Send(message)
}

func main() {
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Println(notify(consoleNotifier{}, "package boundaries stay replaceable"))
}