package main

import (
	"fmt"
	"net/http"
)

func handlerAuthServiceHttp(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Auth Service Http")
}

func main() {
	h := http.HandlerFunc(handlerAuthServiceHttp)
	fmt.Printf("handler=%T topic=Auth Service Http\n", h)
}
