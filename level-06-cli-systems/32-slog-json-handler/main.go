package main

import (
	"fmt"
	"net/http"
)

func handlerSlogJsonHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Slog Json Handler")
}

func main() {
	h := http.HandlerFunc(handlerSlogJsonHandler)
	fmt.Printf("handler=%T topic=Slog Json Handler\n", h)
}
