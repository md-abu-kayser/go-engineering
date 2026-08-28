// Lesson 43: Environment Basics
//
// Goal: Read, check for, and set environment variables from Go code —
// the three operations every CLI tool or service needs for
// configuration via the environment.
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("=== Environment Basics ===")
	fmt.Println("----------------------------------")

	// os.Getenv returns "" if the variable isn't set — it can't tell you
	// the difference between "unset" and "set to an empty string".
	home := os.Getenv("HOME")
	fmt.Printf("os.Getenv(\"HOME\")           : %q\n", home)

	// os.LookupEnv CAN tell the difference, via its second return value.
	value, ok := os.LookupEnv("APP_MODE")
	if !ok {
		fmt.Println(`os.LookupEnv("APP_MODE")     : not set`)
	} else {
		fmt.Printf("os.LookupEnv(\"APP_MODE\")     : %q\n", value)
	}

	// os.Setenv sets a variable for THIS PROCESS ONLY — it does not
	// persist after the program exits, and does not affect your shell.
	os.Setenv("APP_MODE", "debug")
	value, _ = os.LookupEnv("APP_MODE")
	fmt.Printf("after os.Setenv, LookupEnv   : %q\n", value)

	fmt.Printf("\nTotal environment variables visible to this process: %d\n", len(os.Environ()))
}
