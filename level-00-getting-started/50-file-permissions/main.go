// Lesson 50: File Permissions
//
// Goal: Understand Go's os.FileMode, create a file with specific
// permission bits, inspect them, and change them with os.Chmod.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== File Permissions ===")
	fmt.Println("----------------------------------")

	path := "/tmp/lesson50-demo.txt"
	defer os.Remove(path) // clean up after ourselves

	// The permission bits (0644) are the THIRD argument: owner can
	// read+write, group and others can only read. This is the standard
	// Unix permission model, expressed in octal.
	if err := os.WriteFile(path, []byte("hello"), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "error writing file:", err)
		os.Exit(1)
	}

	info, err := os.Stat(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error statting file:", err)
		os.Exit(1)
	}
	fmt.Printf("Created with mode : %s (%#o)\n", info.Mode(), info.Mode().Perm())

	// Tighten permissions to owner-only read+write.
	if err := os.Chmod(path, 0600); err != nil {
		fmt.Fprintln(os.Stderr, "error changing permissions:", err)
		os.Exit(1)
	}

	info, _ = os.Stat(path)
	fmt.Printf("After os.Chmod    : %s (%#o)\n", info.Mode(), info.Mode().Perm())

	fmt.Printf("Is a directory?   : %t\n", info.IsDir())
}
