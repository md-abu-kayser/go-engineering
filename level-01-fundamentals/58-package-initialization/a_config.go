// a_config.go — named to sort alphabetically BEFORE b_database.go, to
// make the file-processing-order point below concrete and checkable.
package main

import "fmt"

// Package-level variable initializers run BEFORE any init() function,
// in dependency order (a variable that reads another variable is
// initialized after the one it depends on) — appVersion has no
// dependencies, so its order relative to other independent variables
// is simply file order, then declaration order within a file.
var appVersion = "v1.0.0"

// init functions run automatically, with NO explicit call anywhere —
// simply defining one is enough. A package can have MULTIPLE init
// functions, even across different files.
func init() {
	fmt.Println("1. init() in a_config.go running (appVersion already initialized:", appVersion, ")")
}
