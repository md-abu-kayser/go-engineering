package main

import (
	"fmt"
	"net/http"
)

func handlerHttptestServer(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Httptest Server")
}

func main() {
	h := http.HandlerFunc(handlerHttptestServer)
	fmt.Printf("handler=%T topic=Httptest Server\n", h)
}
