package main

import (
	"fmt"
	"net/http"
)

func handlerHttpClient(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Client")
}

func main() {
	h := http.HandlerFunc(handlerHttpClient)
	fmt.Printf("handler=%T topic=Http Client\n", h)
}
