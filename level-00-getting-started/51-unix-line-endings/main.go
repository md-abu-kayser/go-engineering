// Lesson 51: Unix Line Endings
//
// Goal: Understand the LF vs CRLF line-ending difference, see how it
// shows up as a literal byte in Go strings, and handle both robustly
// when reading text.
package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("=== Unix Line Endings ===")
	fmt.Println("----------------------------------")

	unixLine := "hello\n"      // LF only  — the Unix/Linux/macOS convention
	windowsLine := "hello\r\n" // CRLF     — the Windows convention

	fmt.Printf("Unix line    : %q (len %d)\n", unixLine, len(unixLine))
	fmt.Printf("Windows line : %q (len %d)\n", windowsLine, len(windowsLine))

	// bufio.Scanner's default line-splitting handles BOTH transparently
	// — it strips a trailing \r\n OR a bare \n, so your code doesn't
	// need to special-case either one.
	mixed := "unix line\nwindows line\r\nanother unix line\n"
	scanner := bufio.NewScanner(strings.NewReader(mixed))
	fmt.Println("\nScanning mixed line endings with bufio.Scanner:")
	for scanner.Scan() {
		fmt.Printf("  %q (no leftover \\r or \\n)\n", scanner.Text())
	}
}
