package main

import (
	"fmt"
	"net/http"
)

func handlerApiDirectory(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Directory")
}

func main() {
	h := http.HandlerFunc(handlerApiDirectory)
	fmt.Printf("handler=%T topic=Api Directory\n", h)
}
