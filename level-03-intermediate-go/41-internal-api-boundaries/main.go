package main

import (
	"fmt"
	"net/http"
)

func handlerInternalApiBoundaries(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Internal Api Boundaries")
}

func main() {
	h := http.HandlerFunc(handlerInternalApiBoundaries)
	fmt.Printf("handler=%T topic=Internal Api Boundaries\n", h)
}
