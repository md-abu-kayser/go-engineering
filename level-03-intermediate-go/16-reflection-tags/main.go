// Lesson 16: Reflection Tags
//
// Goal: Inspect a value at runtime while keeping the reflected operation
// small and explicit.
package main

import (
	"fmt"
	"reflect"
)

const lesson = "Reflection Tags"

type profile struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func describe(value any) string {
	typeOfValue := reflect.TypeOf(value)
	valueOfValue := reflect.ValueOf(value)
	return fmt.Sprintf("type=%s kind=%s value=%v", typeOfValue, valueOfValue.Kind(), value)
}

func main() {
	person := profile{Name: "Asha", Age: 28}
	fmt.Printf("=== %s ===\n", lesson)
	fmt.Println(describe(person))
	fmt.Printf("first field tag: %q\n", reflect.TypeOf(person).Field(0).Tag.Get("json"))
}