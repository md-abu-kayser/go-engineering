package main

import (
	"fmt"
	"net/http"
)

func handlerHttp2Awareness(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http2 Awareness")
}

func main() {
	h := http.HandlerFunc(handlerHttp2Awareness)
	fmt.Printf("handler=%T topic=Http2 Awareness\n", h)
}
