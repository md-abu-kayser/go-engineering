// Lesson 44: Command Arguments
//
// Goal: Read raw command-line arguments via os.Args, then see how the
// standard flag package handles the common case of NAMED flags on top
// of that.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Command Arguments ===")
	fmt.Println("----------------------------------")

	fmt.Printf("os.Args (raw)      : %v\n", os.Args)
	fmt.Printf("program name       : %s\n", os.Args[0])
	fmt.Printf("positional args    : %v\n", os.Args[1:])

	// Now the same kind of input, handled with the flag package instead
	// of manually indexing into os.Args.
	name := flag.String("name", "Gopher", "who to greet")
	shout := flag.Bool("shout", false, "uppercase the greeting")
	flag.Parse()

	greeting := fmt.Sprintf("Hello, %s!", *name)
	if *shout {
		greeting = fmt.Sprintf("HELLO, %s!!!", *name)
	}
	fmt.Println("\n" + greeting)
	fmt.Printf("remaining (non-flag) arguments: %v\n", flag.Args())
}
