// Lesson 52: Windows Line Endings
//
// Goal: Look at the CRLF convention specifically from the "writing
// files" side — Go never silently converts line endings for you, in
// either direction — and see how Git's line-ending settings interact
// with that.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("=== Windows Line Endings ===")
	fmt.Println("----------------------------------")

	path := "/tmp/lesson52-demo.txt"
	defer os.Remove(path)

	lines := []string{"first line", "second line", "third line"}

	// Go's os.WriteFile writes EXACTLY the bytes you give it — there is
	// no automatic "text mode" translation the way some languages'
	// standard I/O has historically had. If you want CRLF, you build it
	// yourself.
	crlfContent := strings.Join(lines, "\r\n") + "\r\n"
	if err := os.WriteFile(path, []byte(crlfContent), 0644); err != nil {
		fmt.Fprintln(os.Stderr, "error writing file:", err)
		os.Exit(1)
	}

	raw, _ := os.ReadFile(path)
	fmt.Printf("Bytes written (CRLF)   : %d\n", len(raw))
	fmt.Printf("Bytes if it were LF    : %d\n", len(strings.Join(lines, "\n"))+1)
	fmt.Printf("Contains \\r\\n?          : %t\n", strings.Contains(string(raw), "\r\n"))

	fmt.Println("\nSee the README for how Git's core.autocrlf and .gitattributes interact")
	fmt.Println("with files like this one across contributors on different operating systems.")
}
