package main

import (
	"fmt"
	"net/http"
)

func handlerHttpClientTimeout(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Client Timeout")
}

func main() {
	h := http.HandlerFunc(handlerHttpClientTimeout)
	fmt.Printf("handler=%T topic=Http Client Timeout\n", h)
}
