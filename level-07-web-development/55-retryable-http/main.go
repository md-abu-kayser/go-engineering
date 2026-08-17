package main

import (
	"fmt"
	"net/http"
)

func handlerRetryableHttp(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Retryable Http")
}

func main() {
	h := http.HandlerFunc(handlerRetryableHttp)
	fmt.Printf("handler=%T topic=Retryable Http\n", h)
}
