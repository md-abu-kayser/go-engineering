package main

import (
	"fmt"
	"net/http"
)

func handlerHttpTransport(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Transport")
}

func main() {
	h := http.HandlerFunc(handlerHttpTransport)
	fmt.Printf("handler=%T topic=Http Transport\n", h)
}
