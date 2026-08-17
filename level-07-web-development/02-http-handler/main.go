package main

import (
	"fmt"
	"net/http"
)

func handlerHttpHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Handler")
}

func main() {
	h := http.HandlerFunc(handlerHttpHandler)
	fmt.Printf("handler=%T topic=Http Handler\n", h)
}
