package main

import (
	"fmt"
	"net/http"
)

func handlerPublicApiBoundaries(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Public Api Boundaries")
}

func main() {
	h := http.HandlerFunc(handlerPublicApiBoundaries)
	fmt.Printf("handler=%T topic=Public Api Boundaries\n", h)
}
