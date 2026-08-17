package main

import (
	"errors"
	"fmt"
)

var errExample = errors.New("example failure")

func main() {
	if err := validate(); err != nil {
		fmt.Printf("Standard Error: %v\n", err)
	}
}

func validate() error {
	return errExample
}
