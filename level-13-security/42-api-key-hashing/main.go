package main

import (
	"fmt"
	"net/http"
)

func handlerApiKeyHashing(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Key Hashing")
}

func main() {
	h := http.HandlerFunc(handlerApiKeyHashing)
	fmt.Printf("handler=%T topic=Api Key Hashing\n", h)
}
