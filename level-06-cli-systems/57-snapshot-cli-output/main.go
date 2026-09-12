// Lesson 57: Snapshot Cli Output
//
// Goal: Parse a command option through an isolated FlagSet so the command
// can be exercised without changing the process-wide flag state.
package main

import (
	"flag"
	"fmt"
	"io"
)

const lesson = "Snapshot Cli Output"

func parse(args []string) (string, error) {
	set := flag.NewFlagSet("lesson", flag.ContinueOnError)
	set.SetOutput(io.Discard)
	name := set.String("name", "engineer", "name to greet")
	if err := set.Parse(args); err != nil {
		return "", err
	}
	return *name, nil
}

func main() {
	name, err := parse([]string{"-name", "Asha"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("hello, %s\n", name)
}