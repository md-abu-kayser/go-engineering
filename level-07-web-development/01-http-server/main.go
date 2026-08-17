package main

import (
	"fmt"
	"net/http"
)

func handlerHttpServer(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Server")
}

func main() {
	h := http.HandlerFunc(handlerHttpServer)
	fmt.Printf("handler=%T topic=Http Server\n", h)
}
