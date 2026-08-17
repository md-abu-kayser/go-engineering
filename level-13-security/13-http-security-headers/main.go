package main

import (
	"fmt"
	"net/http"
)

func handlerHttpSecurityHeaders(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Security Headers")
}

func main() {
	h := http.HandlerFunc(handlerHttpSecurityHeaders)
	fmt.Printf("handler=%T topic=Http Security Headers\n", h)
}
