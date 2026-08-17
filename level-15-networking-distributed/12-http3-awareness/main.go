package main

import (
	"fmt"
	"net/http"
)

func handlerHttp3Awareness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http3 Awareness")
}

func main() {
	h := http.HandlerFunc(handlerHttp3Awareness)
	fmt.Printf("handler=%T topic=Http3 Awareness\n", h)
}
