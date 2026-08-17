package main

import (
	"errors"
	"fmt"
)

var errExample = errors.New("example failure")

func main() {
	if err := validate(); err != nil {
		fmt.Printf("Error Taxonomy: %v\n", err)
	}
}

func validate() error {
	return errExample
}
