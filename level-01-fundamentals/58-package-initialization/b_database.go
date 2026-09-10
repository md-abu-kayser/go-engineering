// b_database.go — sorts alphabetically AFTER a_config.go.
package main

import "fmt"

// connectionString DEPENDS ON appVersion (declared in a_config.go, a
// DIFFERENT file) — Go's initializer ordering handles this correctly
// regardless of file order, because it's based on actual dependencies
// between variables, not just alphabetical file order.
var connectionString = "app-" + appVersion + ".db"

func init() {
	fmt.Println("2. init() in b_database.go running (connectionString:", connectionString, ")")
}

// A SECOND init() function, in the SAME file — perfectly legal. Go
// allows multiple init functions per file too, not just per package.
func init() {
	fmt.Println("3. a SECOND init() in b_database.go — also runs, in the order it appears")
}
