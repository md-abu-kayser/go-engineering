// Command greeter is a second, independent binary living under cmd/,
// demonstrating the cmd/ directory convention: each subdirectory of
// cmd/ is its own separate `package main`, building to its own
// executable, all within one module.
package main

import "fmt"

func main() {
	fmt.Println("I'm a separate binary, built from ./cmd/greeter.")
}
