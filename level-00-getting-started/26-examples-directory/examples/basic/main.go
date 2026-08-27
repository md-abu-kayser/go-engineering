// This is a runnable USAGE EXAMPLE for the greetlib package — not a test,
// and not part of the library itself. It exists purely to show a new
// user of greetlib exactly how to call it, in working, copy-pasteable
// code, which is exactly what an examples/ directory is for.
package main

import (
	"fmt"

	"go-engineering/level-00-getting-started/26-examples-directory/greetlib"
)

func main() {
	defaultGreeting := greetlib.Greeting{}
	fmt.Println(defaultGreeting.For("Gopher"))

	custom := greetlib.Greeting{Prefix: "Welcome"}
	fmt.Println(custom.For("Gopher"))
}
