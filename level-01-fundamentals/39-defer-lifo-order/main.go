// Lesson 39: defer LIFO Order
//
// Goal: See directly that MULTIPLE deferred calls run in LIFO order —
// Last In, First Out — like a stack, and understand why this is
// exactly the right order for nested resource cleanup.
package main

import "fmt"

func main() {
	fmt.Println("=== defer LIFO Order ===")
	fmt.Println("----------------------------------")

	// Three separate defers, registered in order 1, 2, 3 — but they
	// run in the OPPOSITE order: 3, 2, 1. Think of defer as pushing
	// onto a stack; when the function returns, the stack is popped
	// from the top.
	fmt.Println("Registering three deferred calls, in this order: 1, 2, 3...")
	defer fmt.Println("deferred call #1 (registered FIRST, runs LAST)")
	defer fmt.Println("deferred call #2 (registered SECOND, runs SECOND-to-last)")
	defer fmt.Println("deferred call #3 (registered LAST, runs FIRST)")
	fmt.Println("...main() body continues normally here...")
	fmt.Println("...and now main() is about to return; watch the order below:")
}
