package main

import (
	"fmt"
	"net/http"
)

func handlerApiContracts(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprint(w, "Api Contracts")
}

func main() {
	h := http.HandlerFunc(handlerApiContracts)
	fmt.Printf("handler=%T topic=Api Contracts\n", h)
}
