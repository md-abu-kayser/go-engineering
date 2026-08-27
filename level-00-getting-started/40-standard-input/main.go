// Lesson 40: Standard Input
//
// Goal: Read from stdin using bufio.Scanner — the idiomatic way to read
// line-by-line input in Go — and handle EOF gracefully.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Standard Input ===")
	fmt.Println("----------------------------------")
	fmt.Println("Reading lines from stdin until EOF. Try:")
	fmt.Println(`  printf "hello\nworld\n" | go run main.go`)
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		fmt.Printf("line %d (%d chars, uppercased): %s\n", lineNum, len(line), strings.ToUpper(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading stdin:", err)
		os.Exit(1)
	}

	fmt.Printf("\nRead %d line(s) total.\n", lineNum)
}
