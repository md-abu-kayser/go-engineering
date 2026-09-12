// Lesson 53: Rate Limit Abuse
//
// Goal: Validate untrusted input at a boundary and compare authentication
// material in constant time rather than relying on ordinary string equality.
package main

import (
	"crypto/subtle"
	"fmt"
	"strings"
)

const lesson = "Rate Limit Abuse"

func validToken(token string) bool {
	if len(token) != 12 || strings.ContainsAny(token, " \t\n") {
		return false
	}
	expected := "safe-token-1"
	return subtle.ConstantTimeCompare([]byte(token), []byte(expected)) == 1
}

func main() {
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Printf("accepted: %t\n", validToken("safe-token-1"))
	fmt.Printf("rejected: %t\n", validToken("not-a-token!"))
}