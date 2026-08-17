package main

import (
	"fmt"
	"net/http"
)

func handlerApiKeyRotation(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Key Rotation")
}

func main() {
	h := http.HandlerFunc(handlerApiKeyRotation)
	fmt.Printf("handler=%T topic=Api Key Rotation\n", h)
}
