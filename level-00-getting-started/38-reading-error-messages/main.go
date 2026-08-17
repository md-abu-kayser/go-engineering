package main

import (
	"errors"
	"fmt"
)

var errExample = errors.New("example failure")

func main() {
	if err := validate(); err != nil {
		fmt.Printf("Reading Error Messages: %v\n", err)
	}
}

func validate() error {
	return errExample
}
