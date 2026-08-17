package main

import (
	"fmt"
	"net/http"
)

func handlerHttpKeepalive(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Keepalive")
}

func main() {
	h := http.HandlerFunc(handlerHttpKeepalive)
	fmt.Printf("handler=%T topic=Http Keepalive\n", h)
}
