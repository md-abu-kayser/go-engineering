package main

import (
	"fmt"
	"net/http"
)

func handlerApiCompatibilityTest(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Compatibility Test")
}

func main() {
	h := http.HandlerFunc(handlerApiCompatibilityTest)
	fmt.Printf("handler=%T topic=Api Compatibility Test\n", h)
}
