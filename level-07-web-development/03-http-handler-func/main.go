package main

import (
	"fmt"
	"net/http"
)

func handlerHttpHandlerFunc(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Http Handler Func")
}

func main() {
	h := http.HandlerFunc(handlerHttpHandlerFunc)
	fmt.Printf("handler=%T topic=Http Handler Func\n", h)
}
