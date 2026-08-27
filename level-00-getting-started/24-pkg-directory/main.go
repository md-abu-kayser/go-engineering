// Lesson 24: pkg/ Directory
//
// Goal: Use the pkg/ directory convention for exported, reusable library
// code — and understand why the community is split on whether it's a
// good idea.
package main

import (
	"fmt"

	"go-engineering/level-00-getting-started/24-pkg-directory/pkg/stringutil"
)

func main() {
	fmt.Println("=== pkg/ directory ===")
	fmt.Println("----------------------------------")
	fmt.Println(stringutil.Reverse("Gopher"))
	fmt.Println(stringutil.TitleCase("hello from the pkg directory"))
}
